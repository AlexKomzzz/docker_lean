# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /API

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build

EXPOSE 8080

CMD [ "./main" ]