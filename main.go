package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/carldanley/zap2it-scraper/internal/cache"
	"github.com/carldanley/zap2it-scraper/internal/config"
	"github.com/carldanley/zap2it-scraper/internal/providers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if config.ShouldFetchProviders() {
		table, err := providers.FetchTable()
		if err != nil {
			panic(err)
		}

		fmt.Println(table)
	}

	guideCache := cache.New()
	go guideCache.Start()

	http.HandleFunc("/xmltv", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename=\"guide.xmltv\"")
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")

		output, err := guideCache.GetTVGuide().Render()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		_, err = io.WriteString(w, output)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	serverPort := fmt.Sprintf(":%d", config.GetServerPort())
	fmt.Printf("Starting server on port %s\n", serverPort)

	err := http.ListenAndServe(serverPort, nil)
	if errors.Is(err, http.ErrServerClosed) {
		panic("server closed")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}
