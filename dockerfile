FROM golang:latest as build

WORKDIR /build
COPY go.mod .
COPY *.go .
RUN go mod tidy
RUN go build -o csvtodb .
COPY /build/csvtodb .
