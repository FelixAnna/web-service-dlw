resource "azurerm_storage_account" "static_storage" {
  name                     = var.storageAccName
  resource_group_name      = azurerm_resource_group.dlwrg.name
  location                 = azurerm_resource_group.dlwrg.location

  account_kind             = "StorageV2"
  account_tier             = "Standard"
  account_replication_type = "LRS"
  enable_https_traffic_only = true
  min_tls_version = "TLS1_2"

  static_website {
    index_document = "index.html"
  }

  tags = var.tags
}

resource "azurerm_cdn_profile" "dlw_cnd_profile" {
  name                = var.cdnProfileName
  location            = azurerm_resource_group.dlwrg.location
  resource_group_name = azurerm_resource_group.dlwrg.name
  sku                 = "Standard_Microsoft"
}

resource "azurerm_cdn_endpoint" "dlw_origin" {
  name                = var.cdnEndpointName
  profile_name        = azurerm_cdn_profile.dlw_cnd_profile.name
  location            = azurerm_resource_group.dlwrg.location
  resource_group_name = azurerm_resource_group.dlwrg.name
  origin_host_header            = azurerm_storage_account.static_storage.primary_web_host

  delivery_rule {
    name = "rewrite"
    order = 1
    url_file_extension_condition {
      operator = "LessThan"
      match_values = ["1"]
    }
    
    url_rewrite_action {
      source_pattern = "/"
      destination = "/index.html"
      preserve_unmatched_path = false
    }
  }
  origin {
    name      = "dlw"
    host_name = azurerm_storage_account.static_storage.primary_web_host
  }
}

resource "azurerm_cdn_endpoint_custom_domain" "dlw_dns" {
  name            = "dlw-dns"
  cdn_endpoint_id = azurerm_cdn_endpoint.dlw_origin.id
  host_name       = var.frontendDNS

  depends_on = [
    azurerm_cdn_endpoint.dlw_cnd_profile,
    azurerm_cdn_endpoint.dlw_origin
  ]
}
