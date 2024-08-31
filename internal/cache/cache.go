package cache

import (
	"fmt"
	"time"

	"github.com/carldanley/zap2it-scraper/internal/config"
	"github.com/carldanley/zap2it-scraper/pkg/xmltv"
	"github.com/carldanley/zap2it-scraper/pkg/zap2it"
)

const (
	CacheExpirationHours = 3
)

type Cache struct {
	lastFetchTime time.Time
	tvGuide       *xmltv.TVGuide
}

func New() *Cache {
	return &Cache{
		tvGuide: xmltv.CreateTVGuide(config.GetLanguage()),
	}
}

func (c *Cache) IsStale() bool {
	return time.Since(c.lastFetchTime).Hours() > CacheExpirationHours
}

func (c *Cache) Start() {
	fmt.Println("Starting cache engine...")

	for {
		if c.IsStale() {
			fmt.Println("Cache is stale!")

			if err := c.Update(); err != nil {
				fmt.Printf("error updating cache: %s\n", err)
			}
		}

		time.Sleep(time.Hour)
	}
}

func (c *Cache) Update() error {
	fmt.Println("Updating cache...")

	tvGuide := xmltv.CreateTVGuide(config.GetLanguage())

	// calculate the current half hour offset time from 6 hours ago
	currentTimestamp := time.Now().Unix() - (60 * 60 * 6) // 6 hours ago
	halfHourOffset := currentTimestamp % (60 * 30)
	currentTimestamp -= halfHourOffset

	// calculate the end time stamp
	endTimeStamp := currentTimestamp + (86400 * config.GetDaysToFetch())

	// fetch a token
	tokenResponse, err := zap2it.GetTokenResponse(config.GetUsername(), config.GetPassword())
	if err != nil {
		return fmt.Errorf("error getting token: %w", err)
	}

	// fetch all of the guide data
	for currentTimestamp < endTimeStamp {
		// fetch the guide data for the current timestamp
		fmt.Println("Fetching guide data for timestamp:", currentTimestamp)

		guideResponse, err := c.fetchGuideData(tokenResponse.Token, currentTimestamp)
		if err != nil {
			return fmt.Errorf("error fetching guide data: %w", err)
		}

		// iterate through all of the fetched channels and events, updating the guide
		for _, channel := range guideResponse.Channels {
			tvGuide.AddChannel(channel)

			for _, event := range channel.Events {
				tvGuide.AddEvent(event)
			}
		}

		// add 3 hours to the current timestamp
		currentTimestamp += (60 * 60 * 3)
	}

	// if we got this far, it means we fetched all of the content we need
	// so, update the cache
	fmt.Println("Cache updated!")

	c.tvGuide = tvGuide
	c.lastFetchTime = time.Now()

	return nil
}

func (c *Cache) fetchGuideData(token string, unixTimestamp int64) (zap2it.GuideResponse, error) {
	request := zap2it.GuideRequest{
		Token:         token,
		LineupID:      config.GetLineupID(),
		HeadEndID:     config.GetHeadEndID(),
		Device:        config.GetDevice(),
		CountryCode:   config.GetCountryCode(),
		ZipCode:       config.GetZipCode(),
		UnixTimestamp: unixTimestamp,
	}

	response, err := zap2it.GetGuideResponse(request)
	if err != nil {
		return zap2it.GuideResponse{}, fmt.Errorf("could not get guide response: %w", err)
	}

	return response, nil
}

func (c *Cache) GetTVGuide() *xmltv.TVGuide {
	return c.tvGuide
}
