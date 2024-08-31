package xmltv

// Examples of XMLTV format:
// https://wiki.xmltv.org/index.php/XMLTVFormat

type TVGuide struct {
	XMLName           string `xml:"tv"`
	Language          string `xml:"-"`
	SourceInfoURL     string `xml:"source-info-url,attr"`
	SourceInfoName    string `xml:"source-info-name,attr"`
	GeneratorInfoName string `xml:"generator-info-name,attr"`
	GeneratorInfoURL  string `xml:"generator-info-url,attr"`
	Channels          []*Channel
	Programs          []*Program
}

type Channel struct {
	XMLName      string   `xml:"channel"`
	ID           string   `xml:"id,attr"`
	DisplayNames []string `xml:"display-name"`
	Icon         *Icon    `xml:"icon"`
}

type Icon struct {
	XMLName string `xml:"icon"`
	Source  string `xml:"src,attr,omitempty"`
}

type Program struct {
	XMLName   string `xml:"programme"`
	StartTime string `xml:"start,attr"`
	EndTime   string `xml:"stop,attr"`
	ChannelID string `xml:"channel,attr"`

	Title           *Title
	Description     *Description
	Subtitle        *Subtitle
	Length          *Length
	Thumbnail       string `xml:"thumbnail,omitempty"`
	Icon            *Icon
	URL             string `xml:"url,omitempty"`
	EpisodeNumbers  []*EpisodeNumber
	Categories      []*Category
	Subtitles       *Subtitles
	Rating          *Rating
	NewShow         *NewShow
	FinaleShow      *FinaleShow
	PremiereShow    *PremiereShow
	PreviouslyShown *PreviouslyShown
}

type Title struct {
	XMLName  string `xml:"title"`
	Language string `xml:"lang,attr,omitempty"`
	Value    string `xml:",chardata"`
}

type Description struct {
	XMLName  string `xml:"desc"`
	Language string `xml:"lang,attr,omitempty"`
	Value    string `xml:",chardata"` // default: Unavailable
}

type Subtitle struct {
	XMLName  string `xml:"sub-title"`
	Language string `xml:"lang,attr,omitempty"`
	Value    string `xml:",chardata"`
}

type Length struct {
	XMLName string `xml:"length"`
	Units   string `xml:"units,attr,omitempty"`
	Value   string `xml:",chardata"`
}

type EpisodeNumber struct {
	XMLName string `xml:"episode-num"`
	System  string `xml:"system,attr,omitempty"`
	Value   string `xml:",chardata"`
}

type PreviouslyShown struct {
	XMLName string `xml:"previously-shown"`
}

type Subtitles struct {
	XMLName string `xml:"subtitle"`
	Type    string `xml:"type,attr,omitempty"`
}

type Category struct {
	XMLName  string `xml:"category"`
	Language string `xml:"lang,attr,omitempty"`
	Value    string `xml:",chardata"`
}

type Rating struct {
	XMLName string `xml:"rating"`
	Value   string `xml:",chardata"`
}

type NewShow struct {
	XMLName string `xml:"new"`
}

type FinaleShow struct {
	XMLName string `xml:"finale"`
}

type PremiereShow struct {
	XMLName string `xml:"premiere"`
}
