package providers

import (
	"errors"
	"fmt"

	"github.com/carldanley/zap2it-scraper/internal/config"
	"github.com/carldanley/zap2it-scraper/pkg/zap2it"
	"github.com/jedib0t/go-pretty/v6/table"
)

var ErrNoProvidersFound = errors.New("no providers found")

func FetchTable() (string, error) {
	providersResponse, err := zap2it.GetProvidersResponse(config.GetCountryCode(), config.GetZipCode(), config.GetLanguage())
	if err != nil {
		return "", fmt.Errorf("could not get provider response: %w", err)
	}

	if len(providersResponse.Providers) == 0 {
		return "", ErrNoProvidersFound
	}

	message := fmt.Sprintf("Found %d provider(s):\n", len(providersResponse.Providers))

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Name", "Type", "Device", "LineupID", "HeadEndID", "Location", "PostalCode"})

	for _, provider := range providersResponse.Providers {
		tw.AppendRow(table.Row{
			provider.Name,
			provider.Type,
			provider.Device,
			provider.LineupID,
			provider.HeadEndID,
			provider.Location,
			provider.PostalCode,
		})
	}

	return fmt.Sprintf("%s%s", message, tw.Render()), nil
}
