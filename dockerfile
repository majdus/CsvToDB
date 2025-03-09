FROM golang:1.24.0 AS build

WORKDIR /build
COPY go.mod *.go .
RUN go mod tidy
RUN go build -o csvToDB .

FROM scratch
COPY --from=build /build/csvToDB .
