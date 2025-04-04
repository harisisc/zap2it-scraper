# zap2it-scraper

> A simple Go app to scrape EPG data from zap2it and convert it to XMLTV format.

## Introduction

This application will scrape data from Zap2IT in order to build an XMLTV guide for EPG information. Have a look at the environment section below to make sure things are configured properly.

## Rate Limiting

I couldn't easily find the rate limiting information so, instead, I made the caching bust every 3 hours and it will only fetch 6 hours of data prior to the current timestamp + the amount of days to fetch (configured via the environment variable `ZAP2IT_DAYS_TO_FETCH`)

## Fetching Providers

If you are unsure of how to configure this tool, set the `ZAP2IT_FETCH_PROVIDERS` environment variable to `"true"` and you will see an output of available providers for your configuration.

After finding a provider you're satisfied with, update the other environment variables accordingly.

## Environment Variables

| Variable | Description | |
|----------|-------------| --- |
| `ZAP2IT_USERNAME`  | Your [Zap2IT](https://tvlistings.zap2it.com/) Username | `Optional` |
| `ZAP2IT_PASSWORD` | Your [Zap2IT](https://tvlistings.zap2it.com/) Password | `Optional` |
| `ZAP2IT_SERVER_PORT` | The port this server should host the XMLTV guide on | `Required` `Default: 8080` |
| `ZAP2IT_COUNTRY_CODE` | The Zap2IT country code | `Required` `Default: USA` |
| `ZAP2IT_ZIP_CODE` | The Zap2IT zip code | `Required` |
| `ZAP2IT_LINEUP_ID` | The Zap2IT lineup ID. See the `Fetching Providers` section for more information | `Required` `Default DFT` |
| `ZAP2IT_HEADEND_ID` | The Zap2IT headend ID. See the `Fetching Providers` section for more information | `Required` `Default: lineupId` |
| `ZAP2IT_DEVICE` | The Zap2IT device ID; `-` is the default. See the `Fetching Providers` section for more information | `Required` `Default: '-'` |
| `ZAP2IT_LANGUAGE` | The Zap2IT language your guide data should use | `Required` `Default: en` |
| `ZAP2IT_DAYS_TO_FETCH` | The number of days of guide data to fetch from Zap2IT. | `Required` `Default 4` |
| `ZAP2IT_FETCH_PROVIDERS` | Whether or not to fetch the providers and output them as a table during startup. | `Required` `Default: false` |

## Docker build/run
```
docker build -t zap2it-scraper .

docker run \
  --name zap2it-scraper2 \
  -e ZAP2IT_USERNAME="" \
  -e ZAP2IT_PASSWORD="" \
  -e ZAP2IT_SERVER_PORT="8080" \
  -e ZAP2IT_COUNTRY_CODE="USA" \
  -e ZAP2IT_ZIP_CODE="90210" \
  -e ZAP2IT_LINEUP_ID="DFT" \
  -e ZAP2IT_HEADEND_ID="lineupId" \
  -e ZAP2IT_DEVICE="-" \
  -e ZAP2IT_LANGUAGE="en" \
  -e ZAP2IT_DAYS_TO_FETCH="7" \
  -e ZAP2IT_FETCH_PROVIDERS="false" \
  -p 8080:8080 \
  --rm \
  -d \
  zap2it-scraper
```

## Docker Compose Example
```
services:
  zap2it-scraper:
    container_name: zap2it-scraper
    build: ./src/zap2it-scraper
    environment:
      ZAP2IT_USERNAME: ""
      ZAP2IT_PASSWORD: ""
      ZAP2IT_SERVER_PORT: "8080"
      ZAP2IT_COUNTRY_CODE: "USA"
      ZAP2IT_ZIP_CODE: "90210"
      ZAP2IT_LINEUP_ID: "DFT"
      ZAP2IT_HEADEND_ID: "lineupId"
      ZAP2IT_DEVICE: "-"
      ZAP2IT_LANGUAGE: "en"
      ZAP2IT_DAYS_TO_FETCH: "4"
      ZAP2IT_FETCH_PROVIDERS: "false"
    ports:
      - 8080:8080
```