# Ingress Controller

## Ingress for docker-desktop
following: [install nginx](https://kubernetes.github.io/ingress-nginx/deploy/#docker-desktop) to install ingress for certain env.

#### by kubectl: 
    
```bash
kubectl apply -f ingress_deployment.yml
```

#### by helm: 

```bash
helm upgrade --install ingress-nginx ingress-nginx \
--repo https://kubernetes.github.io/ingress-nginx \
--namespace ingress-nginx --create-namespace
```

## kind ingress

ingress for kind is a little difference, reference [kind](../kind/readme.md#enable-ingress-nginx-skip-kong)

## AKS ingress

ingress for AKS , reference [AKS](../../../readme.md#aks)