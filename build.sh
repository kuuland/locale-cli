#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GOPROXY=https://goproxy.cn,direct go build -o kuu-locale && echo 'Linux: kuu-locale'
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 GOPROXY=https://goproxy.cn,direct go build -o kuu-locale.exe && echo 'Windows: kuu-locale.exe'
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 GOPROXY=https://goproxy.cn,direct go build -o kuu-locale_mac && echo 'Mac: kuu-locale_mac'