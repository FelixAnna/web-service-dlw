# Guide for provide infrastructure in azure with AKS + Application Gateway, then deploy services

## infrastructure
[infrastructure](./azurecli/infrastructure.sh): provide application gateway, azure Kubernetes services, and other needed services, so you have a complete environment with a health check, TLS termination, and TLS redirection enabled by default. This only needs to apply one time.

[services](./azurecli/services.sh): deploy/upgrade microservices to the cluster created by above scripts.


## terraform
[infrastructure](./terraform/) deploy by terraform:

```
terraform init
terraform plan
terraform apply -auto-approve
terraform destroy -auto-approve
```

[services](./readme.md)

```
az aks get-credentials --resource-group devRG --name devCluster
helm upgrade --install dlw ./dlw-chart/ --namespace dlwns --create-namespace --values ./dlw-chart/values_aks_appgw.yaml
```
