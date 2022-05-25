# memo api

## functionalities 
1. save user cared date - memo to db (support Gregorian calendar and Lunar calender)
2. search memo for user
3. call date api to calculate recent memo date (MMdd) for given memo


## Docker Guide

### Build

docker build -t memo-api . 
or
docker build -t memo-api:1.0.0 . 
### Check Image

docker image ls

### Tag

docker image tag memo-api:latest memo-api:1.0.0

### Run (use consul for service registry and discovery)

docker run -d -e AWS_ACCESS_KEY_ID=xyz -e AWS_SECRET_ACCESS_KEY=abc -e AWS_REGION=ap-southeast-1 -e profile=dev  --publish 8282:8282 memo-api:1.0.0

## TODO

1. call a api from another docker container #done