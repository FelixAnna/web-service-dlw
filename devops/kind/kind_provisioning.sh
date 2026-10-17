tag=$1
app=$2  # microservice/deployment name
force=$3


if [ "$tag" == '' ];
then
    tag=latest
fi

if [ "$app" == '' ];
then
    app=dlw
fi

cluster_name="desktop" #"${app}-cluster"

if [ "$force" == '-f' ];
then    
  echo "deleting cluster  ..."
  if kind get clusters | grep -qx "$cluster_name"; then
    kind delete clusters $cluster_name
    kind create cluster --config $app-cluster.yml
  fi
  echo "cluster deleted"
fi

ns="${app}ns"

## define your variables somewhere:
## AWS_ACCESS_KEY_ID=xxx
## AWS_SECRET_ACCESS_KEY=xxx
source d:/code/config.sh
echo $AWS_ACCESS_KEY_ID

echo "install nginx  ..."
echo "(if you need kong, please uninstall nginx, then follow readme.md)"
# Upstream: https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
curl -fsSL https://gitee.com/mirrors/ingress-nginx/raw/main/deploy/static/provider/kind/deploy.yaml | sed 's#registry.k8s.io/#k8s.m.daocloud.io/#g' | kubectl apply -f -

echo "install metrics server ..."
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

echo "wait for nginx controller up before install services ..."
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=900s


echo "load images ..."
# docker login -u $ACR_USER -p $ACR_KEY $ACR

# docker pull $ACR/$app-date-api:$tag
# docker pull $ACR/$app-memo-api:$tag
# docker pull $ACR/$app-finance-api:$tag
# docker pull $ACR/$app-user-api:$tag

# kind load docker-image $ACR/$app-date-api:$tag --namen $cluster_name
# kind load docker-image $ACR/$app-memo-api:$tag --name $cluster_name
# kind load docker-image $ACR/$app-finance-api:$tag --name $cluster_name
# kind load docker-image $ACR/$app-user-api:$tag --name $cluster_name


ACR="yufelix"
docker pull $ACR/$app-date-api:$tag
docker pull $ACR/$app-memo-api:$tag
docker pull $ACR/$app-finance-api:$tag
docker pull $ACR/$app-user-api:$tag

kind load docker-image $ACR/$app-date-api:$tag --name $cluster_name
kind load docker-image $ACR/$app-memo-api:$tag --name $cluster_name
kind load docker-image $ACR/$app-finance-api:$tag --name $cluster_name
kind load docker-image $ACR/$app-user-api:$tag --name $cluster_name


echo "install services ..."
cd ..

sed -i "s/awsKeyIdPlaceHolder/$(echo -n $AWS_ACCESS_KEY_ID | base64)/" ./$app-chart-nossl/values_dev.yaml
sed -i "s/awsSecretKeyPlaceHolder/$(echo -n $AWS_SECRET_ACCESS_KEY | base64)/" ./$app-chart-nossl/values_dev.yaml
sed -i "s/imageVersion/$tag/" ./$app-chart-nossl/values_dev.yaml


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