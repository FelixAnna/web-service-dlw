# kind
To work with kind, you need docker installed first, kind create docker containers and use them as nodes to construct the cluster. You may have multiple control planes to make the cluster high availability.

## install kind
install kind with go: [install kind with go](https://kind.sigs.k8s.io/docs/user/quick-start/#installing-with-go-get--go-install)

## create cluster
create a cluster with work node(s) and control-plane(s), and with port 80 mapping to container port 80 (necessary in windows), so we can access the ingress gateway later

    `kind create cluster --config dlw-cluster.yml`

## enable ingress nginx
kind don't support Loadbalance type service, so please use below nginx for kind:

    `kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml`

## pull image issue

in case pull image is very slow in kind cluster's node, you can pull it to local first and then load to kind's nodes by:

    `kind load docker-image xxx:versionxxx --name yourclustername`
