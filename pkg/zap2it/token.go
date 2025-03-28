package zap2it

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	URLTokenEndpoint = "https://tvlistings.gracenote.com/api/user/login"
)

type TokenResponse struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}

func GetTokenResponse(username, password string) (TokenResponse, error) {
	var tokenResponse TokenResponse

	resp, err := resty.New().R().
		SetBody(map[string]string{
			"emailid":        username,
			"password":       password,
			"isfacebookuser": "false",
			"usertype":       "0",
			"objectid":       "",
		}).
		SetResult(&tokenResponse).
		Post(URLTokenEndpoint)
	if err != nil {
		return TokenResponse{}, fmt.Errorf("could not get token response: %w", err)
	}

	switch resp.StatusCode() {
	case http.StatusForbidden:
		return TokenResponse{}, ErrInvalidCredentials
	case http.StatusInternalServerError:
		return TokenResponse{}, ErrInternalServerError
	}

	return tokenResponse, nil
}
