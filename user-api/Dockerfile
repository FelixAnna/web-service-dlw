# syntax=docker/dockerfile:1

FROM golang:1.18
# RUN apk add build-base

ENV GIN_MODE release

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /docker-user-api

EXPOSE 8181

CMD [ "/docker-user-api" ]