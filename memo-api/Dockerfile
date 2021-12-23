# syntax=docker/dockerfile:1

FROM golang:1.17-alpine
RUN apk add build-base

ENV GIN_MODE release

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /docker-memo-api

EXPOSE 8282

CMD [ "/docker-memo-api" ]