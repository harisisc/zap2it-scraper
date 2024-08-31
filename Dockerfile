FROM alpine

LABEL org.opencontainers.image.source https://github.com/carldanley/zap2it-scraper

RUN apk upgrade --no-cache \
  && apk --no-cache add \
  tzdata zip ca-certificates

WORKDIR /usr/share/zoneinfo
RUN zip -r -0 /zoneinfo.zip .
ENV ZONEINFO /zoneinfo.zip

WORKDIR /
ADD zap2it-scraper /bin/

ENTRYPOINT [ "/bin/zap2it-scraper" ]
