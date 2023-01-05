
env=$1
app=$2


ns="${app}ns"

## deploy services
echo "deploy $app micro services"
cd ../../
helm upgrade --install $app ./$app-chart/ --namespace $ns --create-namespace --values ./$app-chart/values_aks_$env.yaml

echo "done"
