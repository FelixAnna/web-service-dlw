data "azurerm_client_config" "current" {}

resource "azurerm_key_vault" "dlwkv" {
    resource_group_name = azurerm_resource_group.dlwrg.name
    location = azurerm_resource_group.dlwrg.location
    
    name = var.valutName
    enabled_for_disk_encryption = true
    tenant_id                   = data.azurerm_client_config.current.tenant_id
    soft_delete_retention_days  = 7
    purge_protection_enabled    = false

    sku_name = "standard"

    access_policy {
        tenant_id = data.azurerm_client_config.current.tenant_id
        object_id = azurerm_user_assigned_identity.gwIdentity.principal_id

        storage_permissions = [
            "Get",
        ]

        certificate_permissions = [
            "Get",
        ]

        secret_permissions = [
            "Get"
        ]
    }

    access_policy {
        tenant_id = data.azurerm_client_config.current.tenant_id
        object_id = data.azurerm_client_config.current.object_id

        certificate_permissions = [
            "Create",
            "Delete",
            "DeleteIssuers",
            "Get",
            "GetIssuers",
            "Import",
            "List",
            "ListIssuers",
            "ManageContacts",
            "ManageIssuers",
            "Purge",
            "SetIssuers",
            "Update",
        ]

        key_permissions = [
            "Backup",
            "Create",
            "Decrypt",
            "Delete",
            "Encrypt",
            "Get",
            "Import",
            "List",
            "Purge",
            "Recover",
            "Restore",
            "Sign",
            "UnwrapKey",
            "Update",
            "Verify",
            "WrapKey",
        ]

        secret_permissions = [
            "Backup",
            "Delete",
            "Get",
            "List",
            "Purge",
            "Recover",
            "Restore",
            "Set",
        ]
    }
}

resource "azurerm_key_vault_certificate" "sslcert" {
    name = "mycert"
    key_vault_id = azurerm_key_vault.dlwkv.id

    certificate_policy {
      issuer_parameters {
        name = "Self"
      }

      key_properties {
        exportable = true
        key_size   = 2048
        key_type   = "RSA"
        reuse_key  = true
      }

      lifetime_action {
        action {
          action_type = "AutoRenew"
        }

        trigger {
          days_before_expiry = 30
        }
      }

      secret_properties {
        content_type = "application/x-pkcs12"
      }

      x509_certificate_properties {
        # Server Authentication = 1.3.6.1.5.5.7.3.1
        # Client Authentication = 1.3.6.1.5.5.7.3.2
        extended_key_usage = ["1.3.6.1.5.5.7.3.1"]

        key_usage = [
            "cRLSign",
            "dataEncipherment",
            "digitalSignature",
            "keyAgreement",
            "keyCertSign",
            "keyEncipherment",
        ]

        subject_alternative_names {
            dns_names = ["internal.contoso.com", "www.dlw.com"]
        }

        subject            = "CN=dlw.com"
        validity_in_months = 12
      }
    }
}

resource "azurerm_role_assignment" "default" {
    scope = azurerm_user_assigned_identity.gwIdentity.id
    principal_id = azurerm_user_assigned_identity.gwIdentity.principal_id
    role_definition_name = "Managed Identity Operator"
}

resource "azurerm_role_assignment" "assignContributor" {
  scope                = azurerm_application_gateway.appGW.id
  role_definition_name = "Contributor"
  principal_id         = azurerm_user_assigned_identity.gwIdentity.principal_id
}

resource "azurerm_role_assignment" "assignReader" {
  scope                = azurerm_resource_group.dlwrg.id
  role_definition_name = "Reader"
  principal_id         = azurerm_user_assigned_identity.gwIdentity.principal_id
}

resource "azurerm_role_assignment" "assignNWContributor" {
  scope                = azurerm_virtual_network.aksVNet.id
  role_definition_name = "Network Contributor"
  principal_id         = azurerm_user_assigned_identity.aksIdentity.principal_id
}


//TODO ensure this identity used by the AGIC (or the one used by it have permission: Reader + Contributor)
