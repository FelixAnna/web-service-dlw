# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.20.1-alpine3.16 AS build
# RUN apk add build-base

# disable cgo : CGO_ENABLED = 0
ENV GIN_MODE release
ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /docker-user-api

##
## Deploy
##
FROM alpine:3.17.0

#RUN apk update \
#   && rm -rf /var/lib/apt/lists/*

WORKDIR /
COPY --from=build /docker-user-api  /docker-user-api

EXPOSE 8181

# Create a group and user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Tell docker that all future commands should run as the appuser user
USER appuser
CMD ["/docker-user-api"]