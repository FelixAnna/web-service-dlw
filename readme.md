# dlw - daily life web microservices

## APIs
1. user api service
2. memo api service
3. date api service
4. finance api service

## deployments
### kubectl
`deployment/kubernetes/*.yaml`: native kubernetes deployments templements, include microservices and nginx ingress.

### helm
`deployment/kubernetes/dlw-helm-autoscaling`: include autoscaling configurations which only supported by kubectl 1.23.* or above, requires latest docker desktop or minikube.

### helm (no autoscaling configuration)
`deployment/kubernetes/dlw-helm`: no autoscaling configurations in the deployments templements

### metrics
`metrics/*.yaml`: enable metrics server which is necessary for horizontalautoscaler or veticalautoscaler, --kubelet-insecure-tls args is used for local, --metric-resolution can be set to longer if use docker-desktop

### kind
`kind/*.yml`: set up kubernetes cluster by using kind, which can run multiple control panel and work nodes by using docker in local

## prepare 
1. register OAuth Apps in https://github.com/settings/developers
   
   the Authorization callback URL should be： http://localhost/user/oauth2/github/redirect
   
   keep the ClientID and ClientSecret

2. Add parameters in aws parameter store: 
   
   https://ap-southeast-1.console.aws.amazon.com/systems-manager/parameters/?region=ap-southeast-1&tab=Table , 
   
   use KMS customer managed keys if necessary.

3. create Tables in aws DynamoDB:

   dlf.Memos, dlf.Users

4. for doker desktop:
    
    a. start your docker-desktop service, enable kubernetes feature (with wsl 2 enbled together).
    
    b. setup ingress-nginx controller by following: https://kubernetes.github.io/ingress-nginx/deploy/#docker-desktop
        or by kubectl: kubectl apply -f ingress_deployment.yml
        or by helm: 
            helm upgrade --install ingress-nginx ingress-nginx \
            --repo https://kubernetes.github.io/ingress-nginx \
            --namespace ingress-nginx --create-namespace

5. for minikube:
    
    a. start minikube: minikube start
    
    b. minikube addons enable ingress
    
    c. eval $(minikube -p minikube docker-env) ## force to use minikube docker deamon in current shell
    
    d. docker build -t xxx-api .  # build image use minikube docker deamon so it visible to minikube
    
    e. you can use  ./dlw-helm-autoscaling when deploy by helm install command 
    
    f. ssh to minikube container to test the api after installed.

6. for kind:

    reference: /kind/readme.md

## deploy by helm
1. download and unzip helm, add folder to env PATH, following: https://helm.sh/   https://github.com/helm/helm/releases

2. add helm chart repo: https://helm.sh/docs/intro/quickstart/
	```bash
	helm repo add bitnami https://charts.bitnami.com/bitnami
	```

3. cd to deployment\kubernetes\dlw-helm-autoscaling, update the awsKeyId and awsSecretKey to correct value in "values.yaml"
4. cd to deployment\kubernetes folder, run:
	```bash
	helm install dlw ./dlw-helm-autoscaling/ --namespace dlw-dev --create-namespace  --values ./dlw-helm-autoscaling/values_dev.yaml
	```
5. after all resources installed, you can access test api from local browser: http://localhost/date/status
6. update by running:
	```bash
	helm upgrade dlw ./dlw-helm-autoscaling/ --namespace dlw-dev --values ./dlw-helm-autoscaling/values_dev.yaml
	```
7. remove all by running:
	```bash
	helm uninstall dlw -n dlw-dev
	```

## deply by kubectl
1. cd to deployment\kubernetes folder, update the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY to correct value in "namespace_config_secret_dev.yaml"
2. build docker images for the 4 api services, and tag them as xxx-api:1.0.0
3. start deployment from deployment folder:
    ```bash
    kubectl apply -f namespace_config_secret_dev.yaml -n dlw-dev
    kubectl apply -f deployment_dev.yaml -n dlw-dev
    kubectl apply -f ingress_dev.yaml -n dlw-dev
    kubectl apply -f auto_scaler.yaml -n dlw-dev
    ```

4. wait for the ingress resource ready
    ```bash
    kubectl config set-context --current --namespace=dlw-dev

    kubectl describe ingress
    ```
5. now, you can access from local browser: http://localhost/date/status

6. clean resource when you finished the local test:
     ```bash
    kubectl delete -f auto_scaler.yaml -n dlw-dev
    kubectl delete -f ingress_dev.yaml -n dlw-dev
    kubectl delete -f deployment_dev.yaml -n dlw-dev
    kubectl delete -f namespace_config_secret_dev.yaml -n dlw-dev
    ```

## local test
1. start consul service and client:  https://learn.hashicorp.com/tutorials/consul/docker-container-agents?in=consul/docker
     a. start server

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

    b. register a client

        ```bash
        docker run \
        --name=tuer \
        consul agent -node=client-1 -join=172.17.0.2
        ```
2. store your aws credentials in place: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
    1. for local recommand: Shared Credentials File
    2. for docker container not in ecs/eks/ec2: use env passed to container (in yaml or from docker run command)
    3. for docker container in ecs/eks/ec2: use aws role assume
3. start user-api on port 8181
    1. run in local system: 
        ```bash
        cd user-api
        profile=dev go run .
        ```
    2. run from local docker container: 
        ```bash
        docker run -d -e AWS_ACCESS_KEY_ID=xyz -e AWS_SECRET_ACCESS_KEY=abc -e AWS_REGION=ap-southeast-1 -e profile=dev  --publish 8383:8383 date-api:1.0.0
        ```
4. start memo-api on port 8282
5. start date-api on port 8383
6. start finance-api on port 8484
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

## AKS deployment
1. create acr, like: dlwcr
2. push local images to the acr:

    docker tag memo-api:1.0.0 dlwcr.azurecr.io/memo-api:1.0.0

    docker push  dlwcr.azurecr.io/date-api:1.0.0
3. create aks cluster, 1 node is ok, select kubeneters >=1.23
4. connect your local kubectl to aks cluster

	`az aks get-credentials --resource-group dlw-cluste_group --name dlw-cluster`

5. install nginx-controller: https://docs.microsoft.com/en-us/azure/aks/ingress-basic?tabs=azure-cli

	add "--set controller.service.externalTrafficPolicy=Local" for enable access to the dynamic assigned public ip of nginx controller

	```
	NAMESPACE=ingress-basic

	helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
	helm repo update

	helm install ingress-nginx ingress-nginx/ingress-nginx \
	  --create-namespace \
	  --namespace $NAMESPACE \
	  --set controller.service.annotations."service\.beta\.kubernetes\.io/azure-load-balancer-health-probe-request-path"=/healthz \
	  --set controller.service.externalTrafficPolicy=Local
	  ```
6. deploy/upgrade/uninstall by：
	
	```
	helm install dlw ./dlw-helm-autoscaling/ --namespace dlw-dev --create-namespace  --values ./dlw-helm-autoscaling/values_aks.yaml

	helm upgrade dlw ./dlw-helm-autoscaling/ --namespace dlw-dev --values ./dlw-helm-autoscaling/values_aks.yaml --set controller.service.externalTrafficPolicy=Local

	helm uninstall dlw -n dlw-dev
	```

7. user external ip of ingress to access the api services
8. metrics server deployed by azure by default
