# kind
To work with kind, you need docker installed first, kind create docker containers and use them as nodes to construct the cluster. You may have multiple control planes to make the cluster high availability.

## **All-In-One** (provisioning infrastructure + deploy microservice)
```
## create or update if exists
sh kind_provisioning.sh 2.1.1

## force re-create
sh kind_provisioning.sh latest dlw -f
```


*if you still want manual install, please following blow instructions.*

## install kind
install kind with go: [install kind with go](https://kind.sigs.k8s.io/docs/user/quick-start/#installing-with-go-get--go-install)

## create cluster
create a cluster with work node(s) and control-plane(s), and with port 80 mapping to container port 80 (necessary in windows), so we can access the ingress gateway later

```
kind create cluster --config dlw-cluster.yml
```

## enable ingress nginx (skip kong)
kind don't support Loadbalance type service, so please use below nginx for kind:

```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
```

## enable ingress kong (skip nginx) (with 2 kind special changes)
```
kubectl apply -f https://raw.githubusercontent.com/Kong/kubernetes-ingress-controller/master/deploy/single/all-in-one-dbless.yaml

kubectl patch deployment -n kong ingress-kong -p '{"spec":{"template":{"spec":{"containers":[{"name":"proxy","ports":[{"containerPort":8000,"hostPort":80,"name":"proxy","protocol":"TCP"},{"containerPort":8443,"hostPort":43,"name":"proxy-ssl","protocol":"TCP"}]}],"nodeSelector":{"ingress-ready":"true"},"tolerations":[{"key":"node-role.kubernetes.io/control-plane","operator":"Equal","effect":"NoSchedule"},{"key":"node-role.kubernetes.io/master","operator":"Equal","effect":"NoSchedule"}]}}}}'

kubectl patch service -n kong kong-proxy -p '{"spec":{"type":"NodePort"}}'
```

## install metrics

install metrics by apply [metrics](../components/metrics/metrics.yaml)

## pull image issue

    in case pull image is very slow in kind cluster's node, you can pull it to local system first and then load to kind's nodes by:

```
kind load docker-image xxx:versionxxx --name yourclustername
```

helm upgrade --install dlw ./dlw-chart-nossl/ --namespace dlwns --create-namespace --values ./dlw-chart-nossl/values_dev.yaml
