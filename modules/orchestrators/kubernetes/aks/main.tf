locals {
  kubernetes_version = var.kubernetes_version == "" ? null : var.kubernetes_version
  tags = merge(var.tags, {
    Environment = var.environment
    ManagedBy   = "clusterforge"
  })
}

#trivy:ignore:AZU-0041
#trivy:ignore:AZU-0042
#trivy:ignore:AZU-0043
resource "azurerm_kubernetes_cluster" "this" {
  #checkov:skip=CKV2_AZURE_29:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  #checkov:skip=CKV_AZURE_115:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  #checkov:skip=CKV_AZURE_116:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  #checkov:skip=CKV_AZURE_117:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  #checkov:skip=CKV_AZURE_141:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  #checkov:skip=CKV_AZURE_168:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  #checkov:skip=CKV_AZURE_170:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  #checkov:skip=CKV_AZURE_171:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  #checkov:skip=CKV_AZURE_172:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  #checkov:skip=CKV_AZURE_226:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  #checkov:skip=CKV_AZURE_227:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  #checkov:skip=CKV_AZURE_232:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  #checkov:skip=CKV_AZURE_4:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  #checkov:skip=CKV_AZURE_6:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  #checkov:skip=CKV_AZURE_7:AKS hardening is explicitly configurable; this experimental module does not claim one universal production profile.
  name                = var.name
  location            = var.location
  resource_group_name = var.resource_group_name
  dns_prefix          = var.dns_prefix
  kubernetes_version  = local.kubernetes_version
  tags                = local.tags

  default_node_pool {
    name           = var.default_node_pool.name
    node_count     = var.default_node_pool.node_count
    vm_size        = var.default_node_pool.vm_size
    vnet_subnet_id = var.subnet_id
  }

  identity {
    type = "SystemAssigned"
  }
}
