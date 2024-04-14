# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.21-alpine AS build
# RUN apk add build-base

# disable cgo : CGO_ENABLED = 0
ENV GIN_MODE release
ENV CGO_ENABLED=0

WORKDIR /app

COPY . ./
WORKDIR /app/user-api

#RUN go mod download
RUN go build -o /docker-user-api

##
## Deploy
##
FROM alpine:3.18.0

#RUN apk update \
#   && rm -rf /var/lib/apt/lists/*

WORKDIR /
COPY --from=build /docker-user-api  /docker-user-api

EXPOSE 8181

# Create a group and user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Tell docker that all future commands should run as the appuser user
USER appuser
CMD ["/docker-user-api"]