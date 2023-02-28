## uninstall basic 
env=$1  # dev or prod
app=$2  # microservice/deployment name

if [ "$app" == '' ];
then
    app=dlw
fi

## define your variables somewhere:
source d:/code/config.sh
echo $AWS_ACCESS_KEY_ID

cd ./services
sh uninstall.sh $env $app
cd ../  # return to current dir: ./aks_nginx

# there is a known bug in deleting domain binded to cdn, so delete the domain manually first
cd ../../
aws route53 change-resource-record-sets --hosted-zone-id $AWS_HOST_ZONE_ID --change-batch file://./azure-devops/deleteHostZoneRecord.json
cd devops/aks_nginx # return to current dir: ./aks_nginx

## destory infrastructure

cd terraform/profiles/$env
terraform init -reconfigure
terraform destroy --target aws_route53_record.web -auto-approve
terraform destroy -auto-approve
