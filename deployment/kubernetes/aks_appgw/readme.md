# Guid for provide infrastructure in azure with AKS + Application Gateway

## infrastructure
[infrastructure](./infrastructure.sh): provide application gateway, azure Kubernetes services, and other needed services, so you have a complete environment with a health check, TLS termination, and TLS redirection enabled by default. This only needs to apply one time.

## services
[services](./services.sh): deploy/upgrade microservices to the cluster created by above scripts.


## terraform

helm upgrade --install dlw ./dlw-helm-autoscaling/ --namespace dlwns --create-namespace --values ./dlw-helm-autoscaling/values_aks_appgw.yaml
