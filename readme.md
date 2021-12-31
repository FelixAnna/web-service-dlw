# dlw - daily life web microservices

## APIs
1. user api service
2. memo api service
3. date api service

## local test
1. set go env: profile=dev
2. start consul service and client:  https://learn.hashicorp.com/tutorials/consul/docker-container-agents?in=consul/docker
     1. start server

        ```bash
        docker run \
        -d \
        -p 8500:8500 \
        -p 8600:8600/udp \
        --name=badger \
        consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
        
        ```

    check ip of the consul server by exec command inside badger server

    ```bash
    docker exec badger consul members
    ```

    2. register a client

        ```bash
        docker run \
        --name=fox \
        consul agent -node=client-1 -join=172.17.0.2
        ```
3. store your aws credentials in place: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
    1. for local recommand: Shared Credentials File
    2. for docker container not in ecs/eks/ec2: use env
    3. for docker container in ecs/eks/ec2: use aws role assume
4. start user-api on port 8181
    1. run in local system: 
        ```bash
        profile=dev go run .
        ```
    2. run from local docker container: 
        ```bash
        docker run -d -e AWS_ACCESS_KEY_ID=xyz -e AWS_SECRET_ACCESS_KEY=abc -e AWS_REGION=ap-southeast-1 -e profile=dev  --publish 8383:8383 date-api:1.0.0
        ```
5. start memo-api on port 8282
6. start date-api on port 8383
7. test your api works
    1. get github redirectUrl: 
    ```bash
    curl --location --request GET 'http://localhost:8181/oauth2/github/authorize/url'
    ```
    2. copy to browser, and keep the token returned
    3. login to user api with the github token
    ```bash
    curl --location --request GET 'http://localhost:8181/oauth2/github/user?access_code=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'
    ```
    4. add one user memo
    ```bash
    curl --location --request PUT 'http://localhost:8282/memos/' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxNjM4MzYzMDY1MDgxIiwiZW1haWwiOiJ5dWVjbnVAaG90bWFpbC5jb20iLCJleHAiOjE2NDAzMjk1MjB9.uOvsu9mLS95Wc9uWONGR-DZx6WPfGxChrHJ6dPaAsag' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "Subject": "Test Item 8",
        "Description":"Test Description 8",
        "MonthDay":1208,
        "StartYear": 1999,
        "Lunar":true
    }'
    ```
    5. find the memo just added, and see the previous and next memo date distance
    ```bash
    curl --location --request GET 'http://localhost:8282/memos/recent?start=1124&end=1227' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxNjM4MzYzMDY1MDgxIiwiZW1haWwiOiJ5dWVjbnVAaG90bWFpbC5jb20iLCJleHAiOjE2NDAzMjk1MjB9.uOvsu9mLS95Wc9uWONGR-DZx6WPfGxChrHJ6dPaAsag'
    ```

## kubernete test

1. setup ingress-nginx controller by following: https://kubernetes.github.io/ingress-nginx/deploy/#docker-desktop
2. cd to deployment folder, update the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY to correct value
3. start deployment from deployment folder:
    ```bash
    kubectl apply -f deployment_dev.yaml
    kubectl apply -f ingress_dev.yaml
    ```

4. wait for the ingress resource ready
    ```bash
    kubectl config set-context --current --namespace=dlw-dev

    kubectl describe ingress
    ```
5. now, you can access from local browser: http://localhost/date/status

6. clean resource when you finished the local test:
     ```bash
    kubectl delete -f ingress_dev.yaml
    kubectl delete -f deployment_dev.yaml
    ```
