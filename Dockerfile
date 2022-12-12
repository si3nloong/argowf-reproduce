FROM golang:1.19-alpine as golang

ARG path

ADD ${path} /go/app
COPY ./go.mod /go/app/go.mod

WORKDIR  /go/app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o main /go/app/main.go 

FROM alpine:latest

COPY --from=golang /usr/local/go/lib/time/zoneinfo.zip /
COPY --from=golang /go/app/main /app/main

ENV ZONEINFO=/zoneinfo.zip

WORKDIR  /app

ENTRYPOINT ./main