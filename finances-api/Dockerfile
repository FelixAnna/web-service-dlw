# syntax=docker/dockerfile:1

FROM golang:1.18
# RUN apk add build-base

ENV GIN_MODE release

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go mod tidy
RUN go build -o /docker-finance-api

EXPOSE 8484

CMD [ "/docker-finance-api" ]