terraform {
  # need setup backend in azure storage account first
  backend "azurerm" {
    resource_group_name  = "configuration-rg"
    storage_account_name = "configstoragefelix"
    container_name       = "tfstate"
    key                  = "dlw-dev.nginx.tfstate"
  }

  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
      version = "3.20.0"
    }

    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

provider "azurerm" {
  features  {
  }
}

provider "aws" {
  region = "us-east-1"
}

locals {
  environment_name = "dev"
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
