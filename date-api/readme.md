## Docker Guide

### Build

docker build -t date-api . 

### Check Image

docker image ls

### Tag

docker image tag date-api:latest date-api:1.0.0

### Run

docker run -d -e AWS_ACCESS_KEY_ID=xyz -e AWS_SECRET_ACCESS_KEY=abc -e AWS_REGION=ap-southeast-1 --publish 8383:8383 date-api:1.0.0