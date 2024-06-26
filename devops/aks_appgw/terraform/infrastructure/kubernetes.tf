resource "random_pet" "prefix" {}

resource "azurerm_kubernetes_cluster" "dlwCluster" {
  name                = var.clusterName
  location            = azurerm_resource_group.dlwrg.location
  resource_group_name = azurerm_resource_group.dlwrg.name
  tags = var.tags

  dns_prefix          = "dlw-${random_pet.prefix.id}"
  kubernetes_version  = "1.28.5"

  default_node_pool {
    name            = "default"
    max_count = 3
    min_count = 1
    //node_count      = 1
    vm_size         = "Standard_B2s"
    os_disk_size_gb = 30
    enable_auto_scaling = true
    vnet_subnet_id = azurerm_subnet.backend.id
  }

  ingress_application_gateway{
    gateway_id = azurerm_application_gateway.appGW.id
  }

  network_profile{
    network_plugin = "azure"
    load_balancer_sku = "standard"
  }
  
  identity {
    type = "SystemAssigned"
  }

  role_based_access_control_enabled = true
}
