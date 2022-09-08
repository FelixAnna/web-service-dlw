# New-SelfSignedCertificate -certstorelocation Cert:\localMachine\my -dnsname "api.metadlw.com"
# $pwd = ConvertTo-SecureString -String "myTest@Pwd123" -Force -AsPlainText
# Export-PfxCertificate -cert Cert:\localMachine\my\A5CE8D378ED5B664D61E59FF57C5D874DDF1CF35  -FilePath C:\Users\Felix_Yu\Downloads\testCert.pfx -Password $pwd
  # https://docs.microsoft.com/en-us/azure/application-gateway/tutorial-ssl-cli

## 
## purge: az keyvault purge -n $vaultName 

## prepare
echo "preparing"

rgName=dlwRG2
region=eastus
ipName=dlwAppGWIp2
vnetName=appGWVnet2
subnetName=gwSubnet2
appgwName=dlwAppGateway2
clusterName=dlwCluster2
identityName=appgwIdentity2
vaultName=dlwVault2
ns=dlwns2

## purge: keyvault
az group delete --name $rgName --location $region -y
az keyvault purge -n $vaultName 

az group create --name $rgName --location $region

## provide certificate
echo "provisioning keyvalt and ssl cert"

# create user managed identity
az identity create -n $identityName -g $rgName -l $region
identityID=$(az identity list --query "[?resourceGroup=='$rgName'].id" -o tsv)
identityPrincipal=$(az identity list --query "[?resourceGroup=='$rgName'].principalId" -o tsv)

# create Azure key vault and certificate (can done through portal as well)
az keyvault create -n $vaultName -g $rgName -l $region

# assign the identity GET certificate access to Azure Key Vault
az keyvault set-policy \
	-n $vaultName \
	-g $rgName \
	--object-id $identityPrincipal \
	--secret-permissions get

# For each new certificate, create a cert on keyvault and add unversioned secret id to Application Gateway
az keyvault certificate create \
	--vault-name $vaultName \
	-n mycert \
	-p "$(az keyvault certificate get-default-policy)"

# assign AGIC identity to have operator access over AppGw identity
az role assignment create --role "Managed Identity Operator" --assignee $identityPrincipal --scope $identityID



## provisioning application gateway
echo "provisioning application gateway"
versionedSecretId=$(az keyvault certificate show -n mycert --vault-name $vaultName --query "sid" -o tsv)
unversionedSecretId=$(echo $versionedSecretId | cut -d'/' -f-5) # remove the version from the url

az network public-ip create -n $ipName -g $rgName --allocation-method Static --sku Standard

az network vnet create -n $vnetName -g $rgName --address-prefix 10.0.0.0/16 \
  --subnet-name $subnetName --subnet-prefix 10.0.0.0/24 

az network application-gateway create -n $appgwName -l $region -g $rgName --sku Standard_v2 \
	--public-ip-address $ipName --vnet-name $vnetName --subnet $subnetName --priority 100 \
	--min-capacity 1 \
	--http-settings-cookie-based-affinity Disabled \
	--frontend-port 443 \
	--http-settings-port 80 \
	--http-settings-protocol Http \
	--identity $identityID \
	--ssl-certificate-name dlwkvsslcert \
	--key-vault-secret-id $unversionedSecretId 
    # ssl certificate with name "mykvsslcert" will be configured on AppGw
    #--cert-file "C:\Users\Felix_Yu\Downloads\testCert.pfx" --cert-password "myTest@Pwd123" 



## provisioning aks
echo "provisioning aks"
appgwId=$(az network application-gateway show -n $appgwName -g $rgName -o tsv --query "id")

az aks create -n $clusterName -g $rgName \
  --kubernetes-version 1.24.3 \
  --vm-set-type VirtualMachineScaleSets --node-count 1 --node-vm-size Standard_B2s \
  --enable-cluster-autoscaler --min-count 1 --max-count 2 \
  --dns-name-prefix dlw \
  --enable-addons ingress-appgw --appgw-id $appgwId \
  --network-plugin azure --enable-managed-identity --generate-ssh-keys



## connect 2 VPC (chance that vnet created, but not found by vnet list command, so this step need manually now)
echo "peering 2 VPCs"

nodeResourceGroup=$(az aks show -n $clusterName -g $rgName -o tsv --query "nodeResourceGroup")
## aksVnetName=$(az network vnet list -g $nodeResourceGroup -o tsv --query "[0].name") ## this command  have bug
aksVnetName=$(az network vnet list --query "[?resourceGroup=='$nodeResourceGroup'].name" --output tsv)
aksVnetId=$(az network vnet show -n $aksVnetName -g $nodeResourceGroup -o tsv --query "id")

az network vnet peering create -n AppGWtoAKSVnetPeering -g $rgName --vnet-name $vnetName \
  --remote-vnet $aksVnetId --allow-vnet-access

appGWVnetId=$(az network vnet show -n $vnetName -g $rgName -o tsv --query "id")

az network vnet peering create -n AKStoAppGWVnetPeering -g $nodeResourceGroup \
  --vnet-name $aksVnetName --remote-vnet $appGWVnetId --allow-vnet-access

groupId=$(az group show -n $rgName -o tsv --query "id")
az role assignment create --role Reader --scope $groupId --assignee $identityPrincipal 
az role assignment create --role Contributor --scope $appgwId --assignee $identityPrincipal
