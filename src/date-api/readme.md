# date api

## functionalities 
1. calculate distance of 2 given date, support Gregorian calendar and Lunar calender
2. return month date info for given date (with leap month, lunar date info)

## Docker Guide

### Build

docker build -t date-api . 
or
docker build -t date-api:1.0.0 . 
### Check Image

docker image ls

### Tag

docker image tag date-api:latest date-api:1.0.0

### Run (use consul for service registry and discovery)

docker run -d -e AWS_ACCESS_KEY_ID=xyz -e AWS_SECRET_ACCESS_KEY=abc -e AWS_REGION=ap-southeast-1 -e profile=dev --publish 8383:8383 date-api:1.0.0