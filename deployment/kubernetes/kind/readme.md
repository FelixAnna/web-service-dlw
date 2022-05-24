## install kind
install kind with go: https://kind.sigs.k8s.io/docs/user/quick-start/#installing-with-go-get--go-install 

## create cluster
create a cluster with one work node and one control-plane, and with port 80 mapping to container port 80 (necessary in windows), so we can access the ingress gateway later
`kind create cluster --config kind-dlw.yml`

## enable ingress nginx
kind don't support Loadbalance type service, so please use below nginx for kind:

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

## pull image issue

in case pull image is very slow, pull it to local first and then load to kind's docker containers by:
`kind load docker-image xxx:versionxxx --name yourclustername`