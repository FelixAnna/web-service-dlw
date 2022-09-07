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

variable "ns" {
    type = string
    description = "(optional) kubernetes namespace to deploy our microservices"
    default = "dlwns"
}

variable "tags" {
    type = map
    description = "(optional) tags for resources"
    default = {
        Application = "dlw"
        Group = "dlwrg"
    }
}
