package xmltv

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/carldanley/zap2it-scraper/pkg/zap2it"
)

func CreateTVGuide(language string) *TVGuide {
	return &TVGuide{
		Language:          language,
		SourceInfoURL:     "https://tvlistings.gracenote.com/",
		SourceInfoName:    "zap2it",
		GeneratorInfoName: "zap2it-scraper",
		GeneratorInfoURL:  "https://github.com/carldanley/zap2it-scraper",
		Channels:          []*Channel{},
		Programs:          []*Program{},
	}
}

func (tvg *TVGuide) ChannelExists(id string) bool {
	for _, channel := range tvg.Channels {
		if channel.ID == id {
			return true
		}
	}

	return false
}

func (tvg *TVGuide) AddChannel(channel zap2it.ChannelResponse) {
	idx := tvg.buildChannelIndex(channel.ChannelNumber, channel.CallSign)

	if tvg.ChannelExists(idx) {
		return
	}

	iconURL := strings.TrimPrefix(channel.Thumbnail, "//")
	iconURL = strings.Split(iconURL, "?")[0]

	newChannel := &Channel{
		ID: idx,
		DisplayNames: []string{
			idx,
			channel.ChannelNumber,
			channel.CallSign,
			strings.ToTitle(channel.AffiliateName),
		},
		Icon: &Icon{
			Source: fmt.Sprintf("http://%s", iconURL),
		},
	}

	tvg.Channels = append(tvg.Channels, newChannel)
}

func (tvg *TVGuide) AddEvent(event zap2it.EventResponse) {
	newProgram := &Program{
		StartTime:      tvg.buildXMLDate(event.StartTime),
		EndTime:        tvg.buildXMLDate(event.EndTime),
		ChannelID:      tvg.buildChannelIndex(event.ChannelNumber, event.CallSign),
		Title:          tvg.buildProgramTitleElement(event),
		Description:    tvg.buildProgramDescriptionElement(event),
		Length:         tvg.buildProgramLengthElement(event),
		URL:            event.GetURL(),
		EpisodeNumbers: tvg.buildProgramEpisodeNumberElements(event),
		Subtitle:       tvg.buildProgramEpisodeSubTitleElement(event),
		Thumbnail:      tvg.buildProgramThumbnailElement(event),
		Icon:           tvg.buildProgramIconElement(event),
		Categories:     tvg.buildProgramCategoryElements(event),
		Subtitles:      tvg.buildProgramSubtitlesElement(event),
		Rating:         tvg.buildProgramRatingElement(event),
	}

	newProgram.processFlags(event)
	tvg.Programs = append(tvg.Programs, newProgram)
}

func (tvg *TVGuide) Render() (string, error) {
	bytes, err := xml.MarshalIndent(tvg, "", "    ")
	if err != nil {
		return "", fmt.Errorf("could not marshal xml: %w", err)
	}

	header := "<?xml version=\"1.0\"?>\n"
	header += "<!DOCTYPE tv SYSTEM \"xmltv.dtd\">\n\n"

	return header + string(bytes), nil
}
