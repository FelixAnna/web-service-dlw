variable "rgName" {
    type = string
    description = "(optional) resource group to deploy your infrastructure"
    default="dlwRG"
}

variable "location" {
    type = string
    description = "(optional) Location"
    default = "eastus"
}

variable "clusterName" {
    type = string
    description = "(optional) azure kubernetes cluster name"
    default = "dlwCluster"
}

variable "ipaddrName" {
    type = string
    description = "(optional) public ip address name for application gateway"
    default = "nginxIp"
}

variable "storageAccName" {
    type = string
    description = "(optional) storage account name for static webhosting"
    default = "dlwstorage916"
}

variable "cdnProfileName" {
    type = string
    description = "(optional) cdn prodile name for static webhosting"
    default = "dlwcdn"
}

variable "cdnEndpointName" {
    type = string
    description = "(optional) cdn endpoint name for static webhosting"
    default = "dlwendpoint"
}

variable "ns" {
    type = string
    description = "(optional) kubernetes namespace to deploy our microservices"
    default = "dlwns"
}

variable "backendDNS" {
    type = string
    description = "(optional) dns record to binding to gateway ipaddress"
    default = "api.metadlw.com"
}

variable "frontendDNS" {
    type = string
    description = "(optional) dns record to binding to gateway ipaddress"
    default = "dev.metadlw.com"
}

variable "tags" {
    type = map
    description = "(optional) tags for resources"
    default = {
        Application = "dlw"
        Group = "dlwrg"
    }
}
