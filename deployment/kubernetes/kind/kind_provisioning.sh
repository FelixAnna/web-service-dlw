kind delete clusters dlw-cluster

kind create cluster --config dlw-cluster.yml

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

kubectl apply -f ../metrics/metrics.yaml

## cd ..
## helm upgrade --install dlw ./dlw-chart/ --namespace dlw-dev --create-namespace --values ./dlw-chart/values_dev.yaml

## kubectl get all -n dlw-dev
