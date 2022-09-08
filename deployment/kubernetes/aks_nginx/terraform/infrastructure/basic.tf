resource "azurerm_resource_group" "dlwrg" {
  name     = var.rgName
  location = var.location
  tags = var.tags
}

resource "azurerm_public_ip" "gwIp" {
  resource_group_name = azurerm_kubernetes_cluster.dlwCluster.node_resource_group
  location            = azurerm_resource_group.dlwrg.location

  name                = var.ipaddrName
  allocation_method   = "Static"
  sku = "Standard"
}
