# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

RUN mkdir -p /common
COPY ../common /common

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./


RUN go build -o /docker-user-api

EXPOSE 8181

CMD [ "/docker-user-api" ]