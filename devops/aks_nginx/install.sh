## provision infrastructure 
app=dlw  # microservice/deployment name
env=$1  # dev or prod

cd ./terraform/profiles/$env
terraform init -reconfigure

terraform apply -auto-approve


## install basic 

cd ../../../services

sh basic_services.sh $app $env

sh main_services.sh $app $env
