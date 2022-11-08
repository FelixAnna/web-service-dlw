# syntax=docker/dockerfile:1

##
## build
##
FROM golang:1.19.3-alpine3.16 AS build
# RUN apk add build-base

# disable cgo : CGO_ENABLED = 0
ENV GIN_MODE release
ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /docker-date-api

##
## Deploy
##
FROM alpine:3.16.0

RUN apk update \
   && rm -rf /var/lib/apt/lists/*

WORKDIR /

COPY --from=build /docker-date-api /docker-date-api

EXPOSE 8383

# Create a group and user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Tell docker that all future commands should run as the appuser user
USER appuser

ENTRYPOINT ["/docker-date-api"]