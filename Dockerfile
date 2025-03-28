FROM golang:alpine AS builder

WORKDIR /src
COPY . .

RUN go build -o zap2it-scraper main.go

FROM alpine

RUN apk upgrade --no-cache \
  && apk --no-cache add \
  tzdata zip ca-certificates

WORKDIR /zap2it-scraper
COPY --from=builder /src/zap2it-scraper .

ENTRYPOINT [ "/zap2it-scraper/zap2it-scraper" ]