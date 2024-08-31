package config

import (
	"os"
	"strconv"
)

const (
	DefaultServerPort  = 8080
	DefaultDaysToFetch = 4
)

func GetUsername() string {
	return os.Getenv("ZAP2IT_USERNAME")
}

func GetPassword() string {
	return os.Getenv("ZAP2IT_PASSWORD")
}

func GetServerPort() int {
	port, exists := os.LookupEnv("ZAP2IT_SERVER_PORT")
	if !exists {
		return DefaultServerPort
	}

	number, err := strconv.Atoi(port)
	if err != nil {
		return DefaultServerPort
	}

	if number == 0 {
		return DefaultServerPort
	}

	return number
}

func GetCountryCode() string {
	code, exists := os.LookupEnv("ZAP2IT_COUNTRY_CODE")
	if !exists {
		return "USA"
	}

	return code
}

func GetZipCode() string {
	return os.Getenv("ZAP2IT_ZIP_CODE")
}

func GetLineupID() string {
	id, exists := os.LookupEnv("ZAP2IT_LINEUP_ID")
	if !exists {
		return "DFT"
	}

	return id
}

func GetHeadEndID() string {
	id, exists := os.LookupEnv("ZAP2IT_HEADEND_ID")
	if !exists {
		return "lineupId"
	}

	return id
}

func GetDevice() string {
	device, exists := os.LookupEnv("ZAP2IT_DEVICE")
	if !exists {
		return "-"
	}

	return device
}

func GetLanguage() string {
	lang, exists := os.LookupEnv("ZAP2IT_LANGUAGE")
	if !exists {
		return "en"
	}

	return lang
}

func GetDaysToFetch() int64 {
	days, exists := os.LookupEnv("ZAP2IT_DAYS_TO_FETCH")
	if !exists {
		return DefaultDaysToFetch
	}

	number, err := strconv.Atoi(days)
	if err != nil {
		return DefaultDaysToFetch
	}

	if number == 0 {
		return DefaultDaysToFetch
	}

	return int64(number)
}

func ShouldFetchProviders() bool {
	fetch, exists := os.LookupEnv("ZAP2IT_FETCH_PROVIDERS")
	if !exists {
		return false
	}

	value, err := strconv.ParseBool(fetch)
	if err != nil {
		return false
	}

	return value
}
