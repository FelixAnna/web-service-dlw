## uninstall basic 
env=$1  # dev or prod
app=$2  # microservice/deployment name

if [ "$app" == '' ];
then
    app=dlw
fi

cd ./services

sh uninstall.sh $env $app

## destory infrastructure

cd ../terraform/profiles/$env
terraform init -reconfigure
terraform destroy --target aws_route53_record.web -auto-approve
terraform destroy -auto-approve
