tag=$1
app=$2  # microservice/deployment name

if [ "$app" == '' ];
then
    app=dlw
fi

if [ "$tag" == '' ];
then
    tag=latest
fi

ns="${app}ns"

## define your variables somewhere:
## AWS_ACCESS_KEY_ID=xxx
## AWS_SECRET_ACCESS_KEY=xxx
source d:/code/config.sh
echo $AWS_ACCESS_KEY_ID

kind delete clusters $app-cluster

kind create cluster --config $app-cluster.yml

echo "install nginx  ..."
echo "(if you need kong, please uninstall nginx, then follow readme.md)"
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

echo "install metrics server ..."
kubectl apply -f ../components/metrics/metrics.yaml

echo "wait for nginx controller up before install services ..."
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=300s

echo "install services ..."
cd ..

sed -i "s/awsKeyIdPlaceHolder/$(echo -n $AWS_ACCESS_KEY_ID | base64)/" ./$app-chart-nossl/values_dev.yaml
sed -i "s/awsSecretKeyPlaceHolder/$(echo -n $AWS_SECRET_ACCESS_KEY | base64)/" ./$app-chart-nossl/values_dev.yaml
sed -i "s/imageVersion/$tag/" ./$app-chart-nossl/values_dev.yaml

# docker login -u $ACR_USER -p $ACR_KEY $ACR

# docker pull $ACR/$app-date-api:$tag
# docker pull $ACR/$app-memo-api:$tag
# docker pull $ACR/$app-finance-api:$tag
# docker pull $ACR/$app-user-api:$tag

# kind load docker-image $ACR/$app-date-api:$tag -n $app-cluster
# kind load docker-image $ACR/$app-memo-api:$tag -n $app-cluster
# kind load docker-image $ACR/$app-finance-api:$tag -n $app-cluster
# kind load docker-image $ACR/$app-user-api:$tag -n $app-cluster

helm upgrade --install $app ./$app-chart-nossl/ \
--namespace $ns \
--create-namespace \
--values ./$app-chart-nossl/values_dev.yaml \
--wait

kubectl get all -n $ns

sed -i "s/$(echo -n $AWS_ACCESS_KEY_ID | base64)/awsKeyIdPlaceHolder/" ./$app-chart-nossl/values_dev.yaml
sed -i "s/$(echo -n $AWS_SECRET_ACCESS_KEY | base64)/awsSecretKeyPlaceHolder/" ./$app-chart-nossl/values_dev.yaml
sed -i "s/$tag/imageVersion/" ./$app-chart-nossl/values_dev.yaml
echo "done"