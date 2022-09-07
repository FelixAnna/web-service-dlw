# New-SelfSignedCertificate -certstorelocation Cert:\localMachine\my -dnsname "www.dlw.com"
# $pwd = ConvertTo-SecureString -String "myTest@Pwd123" -Force -AsPlainText
# Export-PfxCertificate -cert Cert:\localMachine\my\A5CE8D378ED5B664D61E59FF57C5D874DDF1CF35  -FilePath C:\Users\Felix_Yu\Downloads\testCert.pfx -Password $pwd
  # https://docs.microsoft.com/en-us/azure/application-gateway/tutorial-ssl-cli

## 
## purge: az keyvault purge -n $vaultName 

## prepare
echo "preparing"

rgName=dlwRG2
region=eastus
clusterName=dlwCluster2
ns=dlwns2

## purge: rg
az group delete --name $rgName --location $region -y

## provisioning resource group
echo "provisioning resource group"
az group create --name $rgName --location $region

## provisioning aks
echo "provisioning aks"

az aks create -n $clusterName -g $rgName \
  --kubernetes-version 1.24.3 \
  --vm-set-type VirtualMachineScaleSets --node-count 1 --node-vm-size Standard_B2s \
  --enable-cluster-autoscaler --min-count 1 --max-count 2 \
  --dns-name-prefix dlw \
  --network-plugin azure --enable-managed-identity --generate-ssh-keys
