
env=$1
app=$2

ns="${app}ns"

## ensure aks connected
rgName=$app-$env-rg
clusterName="${env}Cluster"
az aks get-credentials --resource-group $rgName --name $clusterName --overwrite-existing

## deploy services
echo "deploy $app micro services"
cd ../../
helm upgrade --install $app ./$app-chart/ --namespace $ns --create-namespace --values ./$app-chart/values_aks_$env.yaml --wait
cd aks_nginx/services ## return to current: ./aks_nginx/services dir

echo "done"
