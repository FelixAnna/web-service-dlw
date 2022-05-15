# deploy prometheus to kubernetes


## deploy

kubectl apply -f configmap.yaml -n dlw-dev
kubectl apply -f role.yaml -n dlw-dev
kubectl apply -f prometheus.yaml -n dlw-dev
kubectl apply -f service.yaml -n dlw-dev

## clean

kubectl delete -f service.yaml -n dlw-dev
kubectl delete -f prometheus.yaml -n dlw-dev
kubectl delete -f role.yaml -n dlw-dev
kubectl delete -f configmap.yaml -n dlw-dev
