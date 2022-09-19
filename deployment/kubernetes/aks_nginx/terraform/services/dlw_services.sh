
ns=dlwns
build="D:\my\github\keep-hands-on\important\dlw-app\build"

## deploy services
echo "deploy dlw micro services"
cd ../../../
helm upgrade --install dlw ./dlw-chart/ --namespace $ns --create-namespace --values ./dlw-chart/values_aks.yaml

echo "deploy dlw frontend"
az storage blob upload-batch --account-name dlwstorage916 -s $build -d '$web' --overwrite

echo "done"
