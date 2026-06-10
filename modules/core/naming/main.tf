locals {
  parts       = [var.project, var.environment, var.name]
  full_name   = join(var.delimiter, local.parts)
  short_name  = substr(local.full_name, 0, var.max_length)
  dns_label   = replace(lower(local.short_name), "_", "-")
  path_prefix = join("/", local.parts)
}
