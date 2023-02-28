
env=$1
app=$2

rgName=$app-$env-rg
ipName=nginxIp
clusterName="${env}Cluster"

## installing basic services
echo "installing basic services ..."

## switch context
az aks get-credentials --resource-group $rgName --name $clusterName --overwrite-existing

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
    --set controller.service.externalTrafficPolicy=Local \
    --wait

echo "wait for nginx controller up before install services ..."
kubectl wait  --namespace $NAMESPACE \
              --for=condition=ready pod \
              --selector=app.kubernetes.io/component=controller \
              --timeout=600s

## deploy consul
echo "deploy consul for service discovery ..."
cd ../../components/consul
sh install.sh
cd ../../aks_nginx/services ## return to current: ./aks_nginx/services dir
# Read more about the installation in the HashiCorp Consul packaged by Bitnami Chart Github repository

echo "installing cert-manager ..."
## config cert-manager
helm repo add jetstack https://charts.jetstack.io
helm repo update

helm upgrade --install cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.10.0 \
  --set installCRDs=true \
  --wait


# or:

# helm repo add bitnami https://charts.bitnami.com/bitnami
# helm install my-release bitnami/consul
# Read more about the installation in the HashiCorp Consul packaged by Bitnami Chart Github repository