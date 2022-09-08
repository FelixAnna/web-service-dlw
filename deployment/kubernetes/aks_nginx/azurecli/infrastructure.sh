# New-SelfSignedCertificate -certstorelocation Cert:\localMachine\my -dnsname "www.dlw.com"
# $pwd = ConvertTo-SecureString -String "myTest@Pwd123" -Force -AsPlainText
# Export-PfxCertificate -cert Cert:\localMachine\my\A5CE8D378ED5B664D61E59FF57C5D874DDF1CF35  -FilePath C:\Users\Felix_Yu\Downloads\testCert.pfx -Password $pwd
  # https://docs.microsoft.com/en-us/azure/application-gateway/tutorial-ssl-cli

## 
## purge: az keyvault purge -n $vaultName 

## prepare
echo "preparing"

rgName=dlwRG2
region=eastus
ipName=nginxIp2
clusterName=dlwCluster2
ns=dlwns2

## purge: rg
az group delete --name $rgName --location $region -y

## provisioning resource group
echo "provisioning resource group"
az group create --name $rgName --location $region


## provisioning aks
echo "provisioning aks"

az aks create -n $clusterName -g $rgName \
  --kubernetes-version 1.24.3 \
  --vm-set-type VirtualMachineScaleSets --node-count 1 --node-vm-size Standard_B2s \
  --enable-cluster-autoscaler --min-count 1 --max-count 2 \
  --dns-name-prefix dlw \
  --network-plugin azure --enable-managed-identity --generate-ssh-keys


## create ip for nginx
nodeResourceGroup=$(az aks show -n $clusterName -g $rgName -o tsv --query "nodeResourceGroup")
az network public-ip create -n $ipName -g $nodeResourceGroup --allocation-method Static --sku Standard


## installing basic services
echo "installing services"

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

echo "done"
