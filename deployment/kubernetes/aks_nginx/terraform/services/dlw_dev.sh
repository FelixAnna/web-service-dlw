
rgName=devrg
ipName=nginxIp
clusterName=devCluster
ns=dlwns

## installing basic services
echo "installing basic services"

## switch context
az aks get-credentials --resource-group $rgName --name $clusterName

## deploy nginx
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

## deploy consul
echo "deploy consul for service discovery and mesh"
cd ../../../../components/consul
sh install.sh

## deploy services
echo "deploy dlw micro services"
cd ../../
helm upgrade --install dlw ./dlw-chart/ --namespace $ns --create-namespace --values ./dlw-chart/values_aks.yaml

echo "done"
