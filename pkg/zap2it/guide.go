package zap2it

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	URLGuideEndpoint = "https://tvlistings.zap2it.com/api/grid"
	URLIconPrefix    = "http://zap2it.tmsimg.com/assets"
)

type GuideRequest struct {
	Token         string `json:"token"`
	LineupID      string `json:"lineupId"`
	HeadEndID     string `json:"headendId"`
	Device        string `json:"device"`
	CountryCode   string `json:"country"`
	ZipCode       string `json:"postalCode"`
	UnixTimestamp int64  `json:"time"`
}

type GuideResponse struct {
	Channels []ChannelResponse `json:"channels"`
}

type ChannelResponse struct {
	CallSign          string          `json:"callSign"`
	AffiliateName     string          `json:"affiliateName"`
	AffiliateCallSign string          `json:"affiliateCallSign"`
	ChannelID         string          `json:"channelId"`
	ChannelNumber     string          `json:"channelNo"`
	Thumbnail         string          `json:"thumbnail"`
	Events            []EventResponse `json:"events"`
}

type EventResponse struct {
	CallSign      string          `json:"callSign"`
	Duration      string          `json:"duration"`
	StartTime     string          `json:"startTime"`
	EndTime       string          `json:"endTime"`
	Thumbnail     string          `json:"thumbnail"`
	ChannelNumber string          `json:"channelNo"`
	SeriesID      string          `json:"seriesId"`
	Rating        string          `json:"rating"`
	Filters       []string        `json:"filter"`
	Flags         []string        `json:"flags"`
	Tags          []string        `json:"tags"`
	Program       ProgramResponse `json:"program"`
}

type ProgramResponse struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	TMSID            string `json:"tmsId"`
	ShortDescription string `json:"shortDescription"`
	Season           string `json:"season"`
	ReleaseYear      string `json:"releaseYear"`
	Episode          string `json:"episode"`
	EpisodeTitle     string `json:"episodeTitle"`
	SeriesID         string `json:"seriesId"`
}

func (e *EventResponse) GetIconURL() string {
	return fmt.Sprintf("%s/%s.jpg", URLIconPrefix, e.Thumbnail)
}

func (e *EventResponse) GetURL() string {
	return fmt.Sprintf("https://tvlistings.zap2it.com/overview.html?programSeriesId=%s&tmsId=%s", e.SeriesID, e.Program.ID)
}

func GetGuideResponse(request GuideRequest) (GuideResponse, error) {
	var guideResponse GuideResponse

	// fetch the response
	resp, err := resty.New().R().
		SetQueryParams(map[string]string{
			// system parameters
			"Activity_ID": "1",
			"FromPage":    "TV Guide",
			"AffiliateId": "gapzap",
			"aid":         "gapzap",
			"timespan":    "3",
			"isOverride":  "true",
			"pref":        "m,p",
			"userId":      "-",

			// user parameters
			"token":      request.Token,
			"lineupId":   request.LineupID,
			"headendId":  request.HeadEndID,
			"device":     request.Device,
			"country":    request.CountryCode,
			"postalCode": request.ZipCode,
			"time":       fmt.Sprintf("%d", request.UnixTimestamp),
		}).
		SetResult(&guideResponse).
		Get(URLGuideEndpoint)
	if err != nil {
		return GuideResponse{}, fmt.Errorf("could not get guide response: %w", err)
	}

	switch resp.StatusCode() {
	case http.StatusBadRequest:
		return GuideResponse{}, ErrBadRequest
	}

	return guideResponse, nil
}
