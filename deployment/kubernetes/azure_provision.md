## provisioning application gateway

  ```bash
	rgName=dlw-rg
	region=eastus
	ipName=dlwAppGWIp
	vnetName=appGWVnet
	subnetName=gwSubnet
	appgwName=dlwAppGateway

	az group create --name $rgName --location $region

	az network public-ip create -n $ipName -g $rgName --allocation-method Static --sku Standard

	az network vnet create -n $vnetName -g $rgName --address-prefix 10.0.0.0/16 \
		--subnet-name $subnetName --subnet-prefix 10.0.0.0/24 

	az network application-gateway create -n $appgwName -l $region -g $rgName --sku Standard_v2 \
		--public-ip-address $ipName --vnet-name $vnetName --subnet $subnetName --priority 100 \
		--min-capacity 1

  ```

## provisioning aks
  refer: [aks](https://docs.microsoft.com/en-us/cli/azure/aks?view=azure-cli-latest#az-aks-create), [application gateway for aks](https://docs.microsoft.com/en-us/azure/application-gateway/tutorial-ingress-controller-add-on-existing#code-try-2)

  * append --no-wait if needed

  ```
  clusterName=dlw-aks
  appgwId=$(az network application-gateway show -n $appgwName -g $rgName -o tsv --query "id") 

  az aks create -n $clusterName -g $rgName \
  --kubernetes-version 1.24.3 \
  --vm-set-type VirtualMachineScaleSets --node-count 1 --node-vm-size Standard_B2s \
  --enable-cluster-autoscaler --min-count 1 --max-count 2 \
  --dns-name-prefix dlw \
  --enable-addons ingress-appgw --appgw-id $appgwId \
  --network-plugin azure --enable-managed-identity --generate-ssh-keys
  ```

## connect 2 VPC

  ```
  nodeResourceGroup=$(az aks show -n $clusterName -g $rgName -o tsv --query "nodeResourceGroup")
  aksVnetName=$(az network vnet list -g $nodeResourceGroup -o tsv --query "[0].name")
  aksVnetId=$(az network vnet show -n $aksVnetName -g $nodeResourceGroup -o tsv --query "id")

  az network vnet peering create -n AppGWtoAKSVnetPeering -g $rgName --vnet-name $vnetName \
    --remote-vnet $aksVnetId --allow-vnet-access

  appGWVnetId=$(az network vnet show -n $vnetName -g $rgName -o tsv --query "id")

  az network vnet peering create -n AKStoAppGWVnetPeering -g $nodeResourceGroup \
    --vnet-name $aksVnetName --remote-vnet $appGWVnetId --allow-vnet-access
  ```
## configure health probs (path=/status)
```
az network application-gateway probe update --gateway-name $appgwName -g $rgName \
	--name pb-dlw-dev-dlw-service-date-8383-dlw-ingress --path /status

az network application-gateway probe update --gateway-name $appgwName -g $rgName \
	--name pb-dlw-dev-dlw-service-finance-8484-dlw-ingress --path /status

az network application-gateway probe update --gateway-name $appgwName -g $rgName \
	--name pb-dlw-dev-dlw-service-memo-8282-dlw-ingress --path /status

az network application-gateway probe update --gateway-name $appgwName -g $rgName \
	--name pb-dlw-dev-dlw-service-user-8181-dlw-ingress --path /status
```

## install services
```
az aks get-credentials --resource-group dlw-rg --name dlw-aks

helm upgrade --install dlw ./dlw-helm-autoscaling/ --namespace dlw-dev --create-namespace --values ./dlw-helm-autoscaling/values_aks_apgw.yaml
```
