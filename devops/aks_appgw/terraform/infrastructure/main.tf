terraform {
  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
      version = "3.99.0"
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
