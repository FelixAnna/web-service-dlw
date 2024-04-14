terraform {
  # need setup backend in azure storage account first
  backend "azurerm" {
    resource_group_name  = "felix-configuration-rg"
    storage_account_name = "configstoragefelix"
    container_name       = "tfstate"
    key                  = "dlw-prod.terraform.tfstate"
  }

  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
      version = "3.99.0"
    }
  }
}

provider "azurerm" {
  features  {
  }
}

locals {
  environment_name = "prod"
}

module "infrastructure" {
  source = "../../infrastructure"

  # Input Variables
  clusterName = "${local.environment_name}Cluster"
  rgName = "dlw-${local.environment_name}-rg"
  backendDNS = "api-${local.environment_name}-dlw.metadlw.com"
  tags = {
      Application = "dlw"
      Group = "dlw"
      Environment="${local.environment_name}"
  }
}
