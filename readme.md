# dlw - daily life web microservices

## APIs
1. user api service
2. memo api service
3. date api service

## local test
1. ensure you have docker service install and started (like docker-desktop)
2. start consul service and client:  https://learn.hashicorp.com/tutorials/consul/docker-container-agents?in=consul/docker
     1. start server

        ```bash
        docker run \
        -d \
        -p 8500:8500 \
        -p 8600:8600/udp \
        --name=shifu \
        consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
        
        ```

    check ip of the consul server by exec command inside badger server

    ```bash
    docker exec badger consul members
    ```

    2. register a client

        ```bash
        docker run \
        --name=tuer \
        consul agent -node=client-1 -join=172.17.0.2
        ```
3. store your aws credentials in place: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
    1. for local recommand: Shared Credentials File
    2. for docker container not in ecs/eks/ec2: use env passed to container (in yaml or from docker run command)
    3. for docker container in ecs/eks/ec2: use aws role assume
4. start user-api on port 8181
    1. run in local system: 
        ```bash
        cd user-api
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

    ```json
    [
        {
            "Id": "1639581432081",
            "UserId": "1638363065081",
            "Subject": "Test Item 7",
            "Description": "Test Description 7",
            "MonthDay": 1208,
            "StartYear": 1999,
            "Lunar": false,
            "Distance": [
                -28,
                337
            ],
            "CreateTime": "1639581432"
        },
        {
            "Id": "1639581443887",
            "UserId": "1638363065081",
            "Subject": "Test Item 8",
            "Description": "Test Description 8",
            "MonthDay": 1208,
            "StartYear": 1999,
            "Lunar": true,
            "Distance": [
                -32,
                323
            ],
            "CreateTime": "1639581443"
        }
    ]
    ```

## kubernete test

1. setup ingress-nginx controller by following: https://kubernetes.github.io/ingress-nginx/deploy/#docker-desktop
2. cd to deployment\kubernetes folder, update the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY to correct value in "namespace_config_secret_dev.yaml"
3. make sure you have build docker images for the 3 api services, and tag them as xxx-api:1.0.0
4. start deployment from deployment folder:
    ```bash
	kubectl apply -f namespace_config_secret_dev.yaml
    kubectl apply -f deployment_dev.yaml
    kubectl apply -f ingress_dev.yaml
    ```

5. wait for the ingress resource ready
    ```bash
    kubectl config set-context --current --namespace=dlw-dev

    kubectl describe ingress
    ```
6. now, you can access from local browser: http://localhost/date/status

7. clean resource when you finished the local test:
     ```bash
    kubectl delete -f ingress_dev.yaml
    kubectl delete -f deployment_dev.yaml
	kubectl delete -f namespace_config_secret_dev.yaml
    ```
## helm test
1. download and unzip helm, add folder to env PATH, following: https://helm.sh/   https://github.com/helm/helm/releases

2. add helm chart repo: https://helm.sh/docs/intro/quickstart/
	```bash
	kubectl create namespace dlw-dev
	
	helm repo add bitnami https://charts.bitnami.com/bitnami
	```
3. cd to deployment\kubernetes\dlw-helm, update the awsKeyId and awsSecretKey to correct value in "values.yaml"
4. cd to deployment\kubernetes folder, run:
	```bash
	helm install dlw ./dlw-helm/
	```
5. after all resources installed, you can access test api from local browser: http://localhost/date/status
6. update by running:
	```bash
	helm upgrade dlw ./dlw-helm/
	```
7. remove all by running:
	```bash
	helm uninstall dlw
	```
## todo
### create user / serviceaccount for containerd service (service registry need rabc, currently use default user) #done at 2022-01-03
### use helm to organize deployment templete #done at 2022-01-04
### 