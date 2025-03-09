FROM golang:latest as build

WORKDIR /build
COPY . .
RUN go mod download
RUN go build  -o csvtodb .

FROM alpine:latest as run
WORKDIR /app
COPY --from=build /build/csvtodb .
COPY /app/data.db .
ENTRYPOINT ["/app/csvtodb"]
