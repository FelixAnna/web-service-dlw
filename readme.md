# dlw - daily life web microservices (RESTful)

[![Build Status](https://github.com/FelixAnna/web-service-dlw/workflows/Run%20Tests/badge.svg?branch=master)](https://github.com/FelixAnna/web-service-dlw/actions?query=branch%3Amaster)
[![codecov](https://codecov.io/gh/FelixAnna/web-service-dlw/branch/master/graph/badge.svg)](https://codecov.io/gh/FelixAnna/web-service-dlw)
[![Go Report Card](https://goreportcard.com/badge/github.com/FelixAnna/web-service-dlw/common)](https://goreportcard.com/report/github.com/FelixAnna/web-service-dlw/common)

# Table of Contents

- 1. [Prepare](#prepare)
- 2. [Switch Context](#swith-kubectl-context)

- 3. [Components](#components)
    - 3.1 [Microservices](#microservices)
    - 3.2 [Ingress controller](#ingress-controller)
    - 3.3 [Metric Server](#metric-server)
    - 3.4 [Dashboard](#dashboard)

- 4. [Deployment](#deployments)
    - 4.1 [Helm](#helm-deployments)

- 5. [Target](#target)
    - 5.1 [Kind](#kind)
    - 5.2 [AKS](#aks)
    
- 6. [Front-end](#front-end)

## Prepare 
1. Register OAuth Apps in https://github.com/settings/developers (2+ for different environment)
   
   the Authorization callback URL should be： {baseApiUrl}/user/oauth2/github/redirect
   
   keep the ClientID and ClientSecret

2. Add parameters in aws parameter store: 
   
   https://ap-southeast-1.console.aws.amazon.com/systems-manager/parameters/?region=ap-southeast-1&tab=Table , 
   
   use KMS customer managed keys if necessary.

3. Create Tables in aws DynamoDB:

   dlf.Memos, dlf.Users
   
4. Prepare an SQL Server instance to store data for finance api, table will be automatic migrated

5. Prepare database in MongoDB atlas (free forever for first 500MB)
	
	* database: dlw_mathematicals
	* collections： answers， questions

## Switch kubectl context

after you connected to aks, you context is attached to aks by default, if you want to check your local kubernetes status, you need switch context:

```bash
kubectl config view
kubectl config use-context kind-dlw-cluster
```

## Components
### Microservices

1. User api service: [user api service](/user-api/readme.md) [![Go Report Card](https://goreportcard.com/badge/github.com/FelixAnna/web-service-dlw/user-api)](https://goreportcard.com/report/github.com/FelixAnna/web-service-dlw/user-api)
2. Memo api service: [memo api service](/memo-api/readme.md) [![Go Report Card](https://goreportcard.com/badge/github.com/FelixAnna/web-service-dlw/memo-api)](https://goreportcard.com/report/github.com/FelixAnna/web-service-dlw/memo-api)
3. Date api service: [date api service](/date-api/readme.md) [![Go Report Card](https://goreportcard.com/badge/github.com/FelixAnna/web-service-dlw/date-api)](https://goreportcard.com/report/github.com/FelixAnna/web-service-dlw/date-api)
4. Finance api service: [finance api service](/finance-api/readme.md) [![Go Report Card](https://goreportcard.com/badge/github.com/FelixAnna/web-service-dlw/finance-api)](https://goreportcard.com/report/github.com/FelixAnna/web-service-dlw/finance-api)

#### helm deployment templetes (autoscaling)

`deployment/kubernetes/dlw-helm-autoscaling`: include autoscaling configurations which only supported by kubectl 1.23+ or above.

### Ingress
reference [ingress](./deployment/kubernetes/ingress/readme.md)


### Metric Server
`deployment/kubernetes/metrics/*.yaml`: enable metrics server which is necessary for horizontalautoscaler or veticalautoscaler if metric server not deployed by default, --kubelet-insecure-tls args is used for local, --metric-resolution can be set to longer if use docker-desktop

cloud based kubernetes already include metric server by default.

### Dashboard
`deployment/kubernetes/dashboard`: follow the instructions to enable dashboard.

## Deployments
### Helm Deployments

#### setup
1. download and unzip helm, add folder to env PATH, following: https://helm.sh/   https://github.com/helm/helm/releases

2. add helm chart repo: https://helm.sh/docs/intro/quickstart/
	```bash
	helm repo add bitnami https://charts.bitnami.com/bitnami
	```
#### deploy
1. update the *awsKeyId* and *awsSecretKey* to correct value in: `deployment\kubernetes\dlw-helm-autoscalingvalues_*.yaml`
2. cd to `deployment\kubernetes` folder, run:
	```bash
	helm install dlw ./dlw-helm-autoscaling/ --namespace dlw-dev --create-namespace  --values ./dlw-helm-autoscaling/values_dev.yaml
	```
3. after all resources installed (include ingress controller), access test api from local browser: http://localhost/date/status
4. update by running:
	```bash
	helm upgrade --install dlw ./dlw-helm-autoscaling/ --namespace dlw-dev --values ./dlw-helm-autoscaling/values_dev.yaml
	```
5. remove all by running:
	```bash
	helm uninstall dlw -n dlw-dev
	```

### AKS
1. create acr, like: dlwcr
2. push local images to the acr like below:

    ```bash
    docker tag memo-api:1.0.0 dlwcr.azurecr.io/memo-api:1.0.0
    docker push  dlwcr.azurecr.io/date-api:1.0.0
    ```

3. create aks cluster, 1 node is ok, select kubeneters >=1.23
4. connect your local kubectl to aks cluster

	`az aks get-credentials --resource-group dlw-cluste_group --name dlw-cluster`

    `az ml computetarget detach -n dlw-cluster -g dlw-cluste_group -w myworkspace`
    `az aks get-credentials --resource-group dlw-cluste_group --name dlw-cluster`

5. install nginx-controller: [install nginx for aks](https://docs.microsoft.com/en-us/azure/aks/ingress-basic?tabs=azure-cli)

	add `--set controller.service.externalTrafficPolicy=Local` for enable access to the dynamic assigned public ip of nginx controller

	```bash
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
	
	```bash
	helm upgrade --install dlw ./dlw-helm-autoscaling/ --namespace dlw-dev --create-namespace --values ./dlw-helm-autoscaling/values_aks.yaml

	helm upgrade dlw ./dlw-helm-autoscaling/ --namespace dlw-dev --values ./dlw-helm-autoscaling/values_aks.yaml --set controller.service.externalTrafficPolicy=Local

	helm uninstall dlw -n dlw-dev
	```

7. user external ip of ingress to access the api services


## Front-end
implemented by [ReactJs + Redux](https://github.com/FelixAnna/keep-hands-on/tree/master/important/dlw-app)
