# memo api

## functionalities 

1. upload file

    example:
    curl --location --request POST 'http://localhost:8484/zdj/upload' \
    --header 'Content-Type:  multipart/form-data' \
    --form 'file=@"/D:/my/github/keep-learning/zdj.txt"'

2. search by criteria
    curl --location --request POST 'http://localhost:8484/zdj/search' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "Distrct":  "福田",
        "Street":  "",
        "Community":  "",

        "MinPrice":  49000,
        "MaxPrice":  50000,

        "Version":  2021,

        "SortKey":  "Price",
        "Page": 1,
        "Size":3
    }'

## Docker Guide

### Build

docker build -t finance-api . 

### Check Image

docker image ls

### Tag

docker image tag finance-api:latest finance-api:1.0.0

### Run (use consul for service registry and discovery)

docker run -d -e AWS_ACCESS_KEY_ID=xyz -e AWS_SECRET_ACCESS_KEY=abc -e AWS_REGION=ap-southeast-1 -e profile=dev  --publish 8484:8484 finance-api:1.0.0