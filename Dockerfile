FROM golang:1.16.0-alpine AS build

RUN apk update && apk add --no-cache git

WORKDIR /build

COPY . .

RUN GOOS=linux go build -ldflags="-w -s" -o mangadex-api

FROM alpine

COPY --from=build /build/mangadex-api /usr/bin/mangadex-api

ENTRYPOINT [ "/usr/bin/mangadex-api" ]
EXPOSE 3000
HEALTHCHECK --interval=30s --timeout=30s \
    CMD curl -f 127.0.0.1:3000/_/healthy || exit 1

