
app=$1
env=$2
build="D:\code\github\keep-hands-on\important\dlw-app\build"
storage=dlwstorage916

ns="${app}ns"

## deploy services
echo "deploy $app micro services"
cd ../../
helm upgrade --install $app ./$app-chart/ --namespace $ns --create-namespace --values ./$app-chart/values_aks_$env.yaml

echo "deploy $app frontend"
az storage blob upload-batch --account-name $storage -s $build -d '$web' --overwrite

echo "done"
