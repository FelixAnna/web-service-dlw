# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.21.1-alpine3.16 AS build
# RUN apk add build-base

# disable cgo : CGO_ENABLED = 0
ENV GIN_MODE release
ENV CGO_ENABLED=0

WORKDIR /app

COPY . ./
WORKDIR /app/memo-api

RUN go build -o /docker-memo-api

##
## Deploy
##
FROM alpine:3.18.0

#RUN apk update \
#   && rm -rf /var/lib/apt/lists/*

WORKDIR /

COPY --from=build /docker-memo-api /docker-memo-api

EXPOSE 8282

# Create a group and user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Tell docker that all future commands should run as the appuser user
USER appuser

ENTRYPOINT ["/docker-memo-api"]