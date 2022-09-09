
rgName=devrg
clusterName=devCluster
ns=dlwns

## installing basic services
echo "removing basic services"

## switch context
az aks get-credentials --resource-group $rgName --name $clusterName

## uninstall services
helm uninstall dlw -n $ns
helm uninstall cert-manager -n cert-manager
helm uninstall ingress-nginx  -n ingress-basic

echo "done"
