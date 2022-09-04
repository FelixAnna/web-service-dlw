terraform {
  # Assumes s3 bucket and dynamo DB table already set up
  # See /code/03-basics/aws-backend
  backend "local" {
    path = "backend/terraform.tfstate"
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
  environment_name = "dev"
}

module "infrastructure" {
  source = "../../dlw"

  # Input Variables
    # Input Variables
  clusterName = "dlw${local.environment_name}Cluster"
  rgName = "dlw${local.environment_name}rg"
  default = {
      Application = "dlw"
      Group = "dlw"
      Environment="${local.environment_name}"
  }
}
