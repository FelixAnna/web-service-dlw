resource "azurerm_resource_group" "dlwrg" {
  name     = var.rgName
  location = var.location
  tags = var.tags
}
