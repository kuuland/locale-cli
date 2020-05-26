#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GOPROXY=https://goproxy.cn,direct go build -o kuu_locale && echo 'Linux: kuu_locale'
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 GOPROXY=https://goproxy.cn,direct go build -o kuu_locale.exe && echo 'Windows: kuu_locale.exe'
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 GOPROXY=https://goproxy.cn,direct go build -o kuu_locale_mac && echo 'Mac: kuu_locale_mac'