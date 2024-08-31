package xmltv

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/carldanley/zap2it-scraper/pkg/zap2it"
)

func (tvg TVGuide) buildXMLDate(time string) string {
	// given format: 20240901040000 +0000
	// expected format: 20240901040000 +0000

	time = strings.Replace(time, "-", "", -1)
	time = strings.Replace(time, "T", "", -1)
	time = strings.Replace(time, ":", "", -1)

	return strings.Replace(time, "Z", " +0000", 1)
}

func (tvg TVGuide) buildChannelIndex(channelNumber, callSign string) string {
	return fmt.Sprintf("%s %s", channelNumber, callSign)
}

func (tvg TVGuide) getProgramIDEpisodeNumber(event zap2it.EventResponse) string {
	programIDLength := len(event.Program.ID)
	seriesID := strings.ReplaceAll(event.SeriesID, "SH", "EP")

	if programIDLength < 4 {
		return fmt.Sprintf("%s.0000", seriesID)
	}

	lastFourCharacters := event.Program.ID[programIDLength-4:]
	return fmt.Sprintf("%s.%s", seriesID, lastFourCharacters)
}

func (tvg TVGuide) buildProgramTitleElement(event zap2it.EventResponse) *Title {
	return &Title{
		Value: event.Program.Title,
	}
}

func (tvg *TVGuide) buildProgramDescriptionElement(event zap2it.EventResponse) *Description {
	description := event.Program.ShortDescription

	if description == "" {
		description = "Unavailable"
	}

	return &Description{
		Value:    description,
		Language: tvg.Language,
	}
}

func (tvg TVGuide) buildProgramLengthElement(event zap2it.EventResponse) *Length {
	return &Length{
		Value: event.Duration,
		Units: "minutes",
	}
}

func (tvg TVGuide) buildProgramEpisodeNumberElements(event zap2it.EventResponse) []*EpisodeNumber {
	episodes := []*EpisodeNumber{
		{
			System: "dd_progid",
			Value:  tvg.getProgramIDEpisodeNumber(event),
		},
	}

	seasonNumber, err := strconv.Atoi(event.Program.Season)
	if err != nil {
		return episodes
	}

	episodeNumber, err := strconv.Atoi(event.Program.Episode)
	if err != nil {
		return episodes
	}

	if seasonNumber > 0 && episodeNumber > 0 {
		episodes = append(episodes, &EpisodeNumber{
			System: "common",
			Value:  fmt.Sprintf("S%02dE%02d", seasonNumber, episodeNumber),
		})

		episodes = append(episodes, &EpisodeNumber{
			Value: fmt.Sprintf("%d.%d", seasonNumber-1, episodeNumber-1),
		})
	}

	return episodes
}

func (tvg *TVGuide) buildProgramEpisodeSubTitleElement(event zap2it.EventResponse) *Subtitle {
	if event.Program.EpisodeTitle != "" {
		return &Subtitle{
			Value:    event.Program.EpisodeTitle,
			Language: tvg.Language,
		}
	}

	return nil
}

func (tvg TVGuide) buildProgramThumbnailElement(event zap2it.EventResponse) string {
	if event.Thumbnail != "" {
		return event.GetIconURL()
	}

	return ""
}

func (tvg TVGuide) buildProgramIconElement(event zap2it.EventResponse) *Icon {
	if event.Thumbnail != "" {
		return &Icon{
			Source: event.GetIconURL(),
		}
	}

	return nil
}

func (tvg *TVGuide) buildProgramCategoryElements(event zap2it.EventResponse) []*Category {
	categories := []*Category{}

	for _, filter := range event.Filters {
		filter = strings.TrimPrefix(filter, "filter-")

		categories = append(categories, &Category{
			Value:    filter,
			Language: tvg.Language,
		})
	}

	seasonNumber, err := strconv.Atoi(event.Program.Season)
	if err != nil {
		return categories
	}

	episodeNumber, err := strconv.Atoi(event.Program.Episode)
	if err != nil {
		return categories
	}

	if seasonNumber > 0 && episodeNumber > 0 {
		categories = append(categories, &Category{
			Value: "Series",
		})
	}

	return categories
}

func (tvg TVGuide) buildProgramSubtitlesElement(event zap2it.EventResponse) *Subtitles {
	for _, tag := range event.Tags {
		if strings.ToLower(tag) == "cc" {
			return &Subtitles{
				Type: "teletext",
			}
		}
	}

	return nil
}

func (tvg TVGuide) buildProgramRatingElement(event zap2it.EventResponse) *Rating {
	if event.Rating != "" {
		return &Rating{
			Value: event.Rating,
		}
	}

	return nil
}

func (p *Program) processFlags(event zap2it.EventResponse) {
	isNewShow := false
	for _, flag := range event.Flags {
		switch strings.ToLower(strings.TrimSpace(flag)) {
		case "new":
			isNewShow = true
			p.NewShow = &NewShow{}
		case "finale":
			p.FinaleShow = &FinaleShow{}
		case "premiere":
			p.PremiereShow = &PremiereShow{}
		}
	}

	if !isNewShow {
		p.PreviouslyShown = &PreviouslyShown{}
	}
}
