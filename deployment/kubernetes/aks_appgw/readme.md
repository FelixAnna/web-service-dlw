# Guid for provide infrastructure in azure with AKS + Application Gateway

## infrastructure
[infrastructure](./infrastructure.sh) provide application gateway, azure kubernetes services and other needed services, so you have a complete environment with health check, TLS termination, TLS redirection enabled by default. This only app

## services
[services](./services.sh) deploy microservices to the cluster created by above scripts
