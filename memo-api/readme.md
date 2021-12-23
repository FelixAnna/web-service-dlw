## Docker Guide

### Build

docker build -t memo-api . 

### Check Image

docker image ls

### Tag

docker image tag memo-api:latest memo-api:1.0.0

### Run

docker run -d -e AWS_ACCESS_KEY_ID=xyz -e AWS_SECRET_ACCESS_KEY=abc -e AWS_REGION=ap-southeast-1 --publish 8282:8282 memo-api:1.0.0

## TODO

1. call a api from another docker container