# Guide for provide infrastructure in azure with AKS with nginx + SSL cert, then deploy services

## azurecli
[infrastructure](./azurecli/infrastructure.sh): provide azure Kubernetes services, ip address (for nginx) and dns bindings(you need bind manually). nginx controller and cert-manager for SSL offloading from nginx, so you have a complete environment with a health check, TLS termination, and TLS redirection enabled by default. This only needs to apply one time.

[services](./azurecli/services.sh): deploy/upgrade microservices to the cluster created by above scripts.


## terraform
[infrastructure](./terraform/) deploy by terraform:

```
terraform init
terraform plan
terraform apply -auto-approve
terraform destroy -auto-approve   # need uninstall nginx helm chart first
```

[configure & install](./terraform/services/) configure environment.
