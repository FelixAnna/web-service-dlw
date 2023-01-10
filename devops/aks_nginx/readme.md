# Guide for provide infrastructure in azure with AKS with nginx + SSL cert, then deploy services

## azurecli
[infrastructure](./azurecli/infrastructure.sh): provide azure Kubernetes services, ip address (for nginx) and dns bindings(you need bind manually). nginx controller and cert-manager for SSL offloading from nginx, so you have a complete environment with a health check, TLS termination, and TLS redirection enabled by default. This only needs to apply one time.

[services](./azurecli/services.sh): deploy/upgrade microservices to the cluster created by above scripts.


## terraform

**All-In-One** install/uninstall:

```
## install everything
sh install.sh dev # dev or prod  

## uninstall everything
sh uninstall.sh dev # dev or prod
```

[infrastructure](./terraform/) apply/destroy by terraform:

```
terraform init
terraform plan
terraform apply -auto-approve

## destroy
terraform destroy -auto-approve   
```
* uninstall infrastructure need uninstall nginx helm chart by using the ./terraform/services/uninstall.sh first

[configure & install](./terraform/services/) configure environment and install services (microservice + frontend).

```
cd services
sh basic_services.sh dev dlw
sh main_services.sh dev dlw

## uninstall before destroy
sh uninstall.sh dev dlw
```

## workaround for cdn custom domain delete issue

delete it manually before uninstall everything
```
aws route53 change-resource-record-sets --hosted-zone-id Z022788838QVNX9DPP3IQ --change-batch '{
  "Comment": "Delete host zone record",
  "Changes": [
    {
      "Action": "DELETE",
      "ResourceRecordSet": {
        "Name": "www.metadlw.com",
        "Type": "CNAME",
		"TTL": 10,
		"ResourceRecords": [
		  {
			"Value": "dlwendpoint.azureedge.net"
		  }]
      }
    }
  ]
}'
```
