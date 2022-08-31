kind delete clusters dlw-cluster

kind create cluster --config dlw-cluster.yml

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

kubectl apply -f ../metrics/metrics.yaml

## cd ..
## helm upgrade --install dlw ./dlw-helm-autoscaling/ --namespace dlw-dev --create-namespace --values ./dlw-helm-autoscaling/values_dev.yaml

## kubectl get all -n dlw-dev
