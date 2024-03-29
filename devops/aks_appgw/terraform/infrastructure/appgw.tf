resource "azurerm_public_ip" "gwIp" {
  resource_group_name = azurerm_resource_group.dlwrg.name
  location            = azurerm_resource_group.dlwrg.location

  name                = var.ipaddrName
  allocation_method   = "Static"
  sku = "Standard"
}

locals {
  gateway_ip_configure_name      = "${azurerm_virtual_network.gwVNet.name}-gwipn"
  backend_address_pool_name      = "${azurerm_virtual_network.gwVNet.name}-beap"
  frontend_port_name             = "${azurerm_virtual_network.gwVNet.name}-feport"
  frontend_ip_configuration_name = "${azurerm_virtual_network.gwVNet.name}-feip"
  http_setting_name              = "${azurerm_virtual_network.gwVNet.name}-be-htst"
  listener_name                  = "${azurerm_virtual_network.gwVNet.name}-httplstn"
  request_routing_rule_name      = "${azurerm_virtual_network.gwVNet.name}-rqrt"
  redirect_configuration_name    = "${azurerm_virtual_network.gwVNet.name}-rdrcfg"
}

resource "azurerm_application_gateway" "appGW" {
  resource_group_name = azurerm_resource_group.dlwrg.name
  location            = azurerm_resource_group.dlwrg.location
  tags = var.tags

  name                = var.appgwName

  sku {
    name     = "Standard_v2"
    tier     = "Standard_v2"
    capacity = 1
  }

  gateway_ip_configuration {
    name      = local.gateway_ip_configure_name
    subnet_id = azurerm_subnet.frontend.id
  }

  frontend_port {
    name = local.frontend_port_name
    port = 443
  }

  frontend_ip_configuration {
    name                 = local.frontend_ip_configuration_name
    public_ip_address_id = azurerm_public_ip.gwIp.id
  }

  backend_address_pool {
    name = local.backend_address_pool_name
  }
  
  identity {
    type = "UserAssigned"
    identity_ids = [azurerm_user_assigned_identity.gwIdentity.id]
  }

  backend_http_settings {
    name                  = local.http_setting_name
    cookie_based_affinity = "Disabled"
    port                  = 80
    protocol              = "Http"
    request_timeout       = 15
  }

  ssl_certificate {
    name = var.sslCertName
    key_vault_secret_id = azurerm_key_vault_certificate.sslcert.secret_id
  }

  http_listener {
    name                           = local.listener_name
    frontend_ip_configuration_name = local.frontend_ip_configuration_name
    frontend_port_name             = local.frontend_port_name
    protocol                       = "Https"
    ssl_certificate_name = var.sslCertName
  }

  request_routing_rule {
    name                       = local.request_routing_rule_name
    rule_type                  = "Basic"
    http_listener_name         = local.listener_name
    backend_address_pool_name  = local.backend_address_pool_name
    backend_http_settings_name = local.http_setting_name
    priority = 100
  }
}
