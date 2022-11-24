## uninstall basic 
app=dlw  # microservice/deployment name
env=$1  # dev or prod

cd ./services

sh uninstall.sh $app $env

## destory infrastructure

cd ../terraform/profiles/$env

terraform destroy -auto-approve
