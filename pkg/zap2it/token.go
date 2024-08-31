package zap2it

import (
	"errors"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	URLTokenEndpoint = "https://tvlistings.zap2it.com/api/user/login"
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
		return TokenResponse{}, err
	}

	switch resp.StatusCode() {
	case http.StatusForbidden:
		return TokenResponse{}, errors.New(ErrInvalidCredentials)
	case http.StatusInternalServerError:
		return TokenResponse{}, errors.New(ErrInternalServerError)
	}

	return tokenResponse, err
}
