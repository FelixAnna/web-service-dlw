resource "azurerm_resource_group" "dlwrg" {
  name     = var.rgName
  location = var.region
  tags = var.tags
}

resource "azurerm_user_assigned_identity" "gwIdentity" {
  resource_group_name = azurerm_resource_group.dlwrg.name
  location = azurerm_resource_group.dlwrg.location

  name = var.identityNameGw
}

resource "azurerm_user_assigned_identity" "aksIdentity" {
  resource_group_name = azurerm_resource_group.dlwrg.name
  location = azurerm_resource_group.dlwrg.location

  name = var.identityNameAks
}


resource "azurerm_virtual_network" "gwVNet" {
  resource_group_name = azurerm_resource_group.dlwrg.name
  location            = azurerm_resource_group.dlwrg.location

  name                = var.vnetName
  address_space       = ["10.1.0.0/16"]
}

resource "azurerm_subnet" "frontend" {
  name                 = var.frontSubnetName
  resource_group_name  = azurerm_resource_group.dlwrg.name
  virtual_network_name = azurerm_virtual_network.gwVNet.name
  address_prefixes     = ["10.1.0.0/24"]
}

resource "azurerm_virtual_network" "aksVNet" {
  resource_group_name = azurerm_resource_group.dlwrg.name
  location            = azurerm_resource_group.dlwrg.location

  name                = "aksVNet"
  address_space       = ["10.2.0.0/16"]
}
resource "azurerm_subnet" "backend" {
  name                 = "backend"
  resource_group_name  = azurerm_resource_group.dlwrg.name
  virtual_network_name = azurerm_virtual_network.aksVNet.name
  address_prefixes     = ["10.2.0.0/22"]
}

resource "azurerm_virtual_network_peering" "gw2aks" {
  name                      = "peergw2aks"
  resource_group_name       = azurerm_resource_group.dlwrg.name
  virtual_network_name      = azurerm_virtual_network.gwVNet.name
  remote_virtual_network_id = azurerm_virtual_network.aksVNet.id
}

resource "azurerm_virtual_network_peering" "aks2gw" {
  name                      = "peeraks2gw"
  resource_group_name       = azurerm_resource_group.dlwrg.name
  virtual_network_name      = azurerm_virtual_network.aksVNet.name
  remote_virtual_network_id = azurerm_virtual_network.gwVNet.id
}
