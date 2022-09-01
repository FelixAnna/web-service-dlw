## install services
echo "installing services"
rgName=dlwRG
clusterName=dlw-dev
ns=dlwns
az aks get-credentials --resource-group $rgName --name $clusterName

helm upgrade --install dlw ./dlw-helm-autoscaling/ --namespace $ns --create-namespace --values ./dlw-helm-autoscaling/values_aks_appgw.yaml

echo "done"
