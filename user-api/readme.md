# user api

## functionalities 

1. user can login by github account

2. user can fire CRUD operations on user table after login /register

3. generate native jwt token for user after login

4. filter by jwt token for some api

* support login from own oauth2 server (enhance later)


## Docker Guide

### Build

docker build -t user-api . 

### Check Image

docker image ls

### Tag

docker image tag user-api:latest user-api:1.0.0

### Run (use consul for service registry and discovery)

docker run -d -e AWS_ACCESS_KEY_ID=xyz -e AWS_SECRET_ACCESS_KEY=abc -e AWS_REGION=ap-southeast-1 -e profile=dev  --publish 8181:8181 user-api:1.0.0