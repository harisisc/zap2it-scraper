package zap2it

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

const (
	URLListingProvidersEndpoint = "https://tvlistings.zap2it.com/gapzap_webapi/api/Providers/getPostalCodeProviders"
)

type ListingProvidersResponse struct {
	DSTUTCOffset string `json:"DSTUTCOffset"`
	STDUTCOffset string `json:"StdUTCOffset"`
	DSTStart     string `json:"DSTStart"`
	DSTEnd       string `json:"DSTEnd"`
	Primetime    string `json:"primetime"`
	Providers    []ProviderResponse
}

type ProviderResponse struct {
	Name              string `json:"name"`
	Type              string `json:"type"`
	Device            string `json:"device"`
	LineupID          string `json:"lineupID"`
	HeadEndID         string `json:"headendID"`
	Location          string `json:"location"`
	Timezone          string `json:"timezone"`
	IsDefaultProvider string `json:"isDefaultProvider"`
	PostalCode        string `json:"postalCode"`
}

func GetProvidersResponse(countryCode, zipCode, language string) (ListingProvidersResponse, error) {
	var listingProvidersResponse ListingProvidersResponse

	// prepare the URL
	url := fmt.Sprintf("%s/%s/%s/gapzap", URLListingProvidersEndpoint, countryCode, zipCode)

	if strings.TrimSpace(language) == "" {
		language = "en"
	}

	url = fmt.Sprintf("%s/%s", url, language)

	// fetch the response
	resp, err := resty.New().R().
		SetResult(&listingProvidersResponse).
		Get(url)
	if err != nil {
		return ListingProvidersResponse{}, err
	}

	switch resp.StatusCode() {
	case http.StatusInternalServerError:
		return ListingProvidersResponse{}, errors.New(ErrInternalServerError)
	}

	return listingProvidersResponse, nil
}
