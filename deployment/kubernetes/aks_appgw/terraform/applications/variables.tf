variable "rgName" {
    type = string
    description = "(optional) resource group to deploy your infrastructure"
    default="dlwRG"
}

variable "region" {
    type = string
    description = "(optional) Location"
    default = "eastus"
}

variable "identityNameGw" {
    type = string
    description = "(optional) name of the user managed identity"
    default = "appgwIdentity"
}

variable "identityNameAks" {
    type = string
    description = "(optional) name of the user managed identity"
    default = "aksIdentity"
}

variable "valutName" {
    type = string
    description = "(optional) name of key vault"
    default = "dlwVault"
}

variable "ipaddrName" {
    type = string
    description = "(optional) public ip address name for application gateway"
    default = "dlwAppGWIp"
}

variable "vnetName" {
    type = string
    description = "(optional) name of virtual network for application gateway"
    default = "appGWVnet"
}

variable "frontSubnetName" {
    type = string
    description = "(optional) subnet name to deploy application gateway"
    default = "gwSubnet"
}

variable "backendSubnetName" {
    type = string
    description = "(optional) subnet name to deploy kubernetes"
    default = "appSubnet"
}

variable "appgwName" {
    type = string
    description = "(optional) name of the applicatoon gateway"
    default = "dlwAppGateway"
}

variable "clusterName" {
    type = string
    description = "(optional) azure kubernetes cluster name"
    default = "dlwCluster"
}

variable "ns" {
    type = string
    description = "(optional) kubernetes namespace to deploy our microservices"
    default = "dlwns"
}

variable "sslCertName" {
    type = string
    description = "(optional) ssl cert name in application gateway http listener"
    default = "dlwkvsslcert"
}
