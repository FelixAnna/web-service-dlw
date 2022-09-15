
helm repo add hashicorp https://helm.releases.hashicorp.com
helm search repo hashicorp/consul

helm install consul hashicorp/consul --create-namespace --namespace consul --values ./config.yaml

# kubectl port-forward service/consul-server --namespace consul 8500:8500
# kubectl port-forward service/consul-server --namespace consul 8501:8501
# kubectl get secrets/consul-bootstrap-acl-token --template='{{.data.token | base64decode }}'
