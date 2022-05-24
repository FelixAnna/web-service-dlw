# deploy prometheus to kubernetes


## helm install kube-prometheus-stack
github: https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack

Installs the kube-prometheus stack, a collection of Kubernetes manifests, Grafana dashboards, and Prometheus rules combined with documentation and scripts to provide easy to operate end-to-end Kubernetes cluster monitoring with Prometheus using the Prometheus Operator.

### set up
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

helm install prometheus prometheus-community/kube-prometheus-stack --namespace prometheus --create-namespace

--patch for bugs
kubectl patch ds prometheus-prometheus-node-exporter --type "json" -p '[{"op": "remove", "path" : "/spec/template/spec/containers/0/volumeMounts/2/mountPropagation"}]' -n prometheus

-- expose service and view dashboards

kubectl port-forward service/prometheus-kube-prometheus-prometheus -n prometheus 9090
kubectl port-forward service/prometheus-grafana -n prometheus 8080:80

find adminPassword: https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml, default is "admin:prom-operator" and login to grafana

### clean
helm uninstall prometheus -n prometheus

## (not recommand)
### self deploy 

kubectl apply -f configmap.yaml -n dlw-dev
kubectl apply -f role.yaml -n dlw-dev
kubectl apply -f prometheus.yaml -n dlw-dev
kubectl apply -f service.yaml -n dlw-dev

### self clean

kubectl delete -f service.yaml -n dlw-dev
kubectl delete -f prometheus.yaml -n dlw-dev
kubectl delete -f role.yaml -n dlw-dev
kubectl delete -f configmap.yaml -n dlw-dev


## reference: 
https://github.com/prometheus/prometheus/blob/release-2.35/documentation/examples/prometheus-kubernetes.yml 