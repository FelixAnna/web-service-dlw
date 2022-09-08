## install services
echo "installing services"

rgName=dlwRG2
clusterName=dlwCluster2
ns=dlwns2
ipName=nginxIp2
## switch context
az aks get-credentials --resource-group $rgName --name $clusterName


nodeResourceGroup=$(az aks show -n $clusterName -g $rgName -o tsv --query "nodeResourceGroup")
STATIC_IP=$(az network public-ip show -n $ipName -g $nodeResourceGroup --query "ipAddress" -o tsv)
NAMESPACE=ingress-basic
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm upgrade --install ingress-nginx ingress-nginx/ingress-nginx \
    --create-namespace \
    --namespace $NAMESPACE \
    --set controller.service.loadBalancerIP=$STATIC_IP \
    --set controller.service.annotations."service\.beta\.kubernetes\.io/azure-load-balancer-health-probe-request-path"=/healthz \
    --set controller.service.externalTrafficPolicy=Local

## config cert-manager
helm repo add jetstack https://charts.jetstack.io
helm repo update

helm install cert-manager jetstack/cert-manager \
  --namespace cert-manager  --create-namespace \
  --version v1.9.1 \
  --set installCRDs=true

## deploy services
helm upgrade --install dlw ./dlw-chart/ --namespace $ns --create-namespace --values ./dlw-chart/values_aks.yaml

echo "done"


## create ssl cert
#HOST=metadlw.com
#KEY_FILE=key.txt
#CERT_FILE=cert.txt
#KEY_FILE=tls-secret
#openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ${KEY_FILE} -out ${CERT_FILE} -subj '/CN=${HOST}/O=${HOST}'
#kubectl create namespace $ns
#kubectl create secret tls ${CERT_NAME} --key ${KEY_FILE} --cert ${CERT_FILE} -n $ns
