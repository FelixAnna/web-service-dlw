terraform {
  # need setup backend in azure storage account first
  backend "azurerm" {
    resource_group_name  = "configuration-rg"
    storage_account_name = "configstoragefelix"
    container_name       = "tfstate"
    key                  = "prod.terraform.tfstate"
  }

  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
      version = "3.20.0"
    }
  }
}

provider "azurerm" {
  features  {
    key_vault {
      purge_soft_delete_on_destroy = true
    }
  }
}

locals {
  environment_name = "prod"
}

module "infrastructure" {
  source = "../../infrastructure"

  # Input Variables
  clusterName = "${local.environment_name}Cluster"
  rgName = "${local.environment_name}rg"
  tags = {
      Application = "dlw"
      Group = "dlw"
      Environment="${local.environment_name}"
  }
}
