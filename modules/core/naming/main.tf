locals {
  raw_parts = concat(
    [
      var.project,
      var.environment,
      var.component,
      var.name
    ],
    var.extra_parts,
    [var.suffix]
  )

  parts       = [for part in local.raw_parts : trimspace(part) if length(trimspace(part)) > 0]
  joined_name = join(var.separator, local.parts)
  cased_name  = var.lowercase ? lower(local.joined_name) : local.joined_name

  supported_pattern = var.separator == "_" ? "/[^0-9A-Za-z_]+/" : (
    var.separator == "-" ? "/[^0-9A-Za-z-]+/" : "/[^0-9A-Za-z]+/"
  )

  normalized_name = replace(local.cased_name, local.supported_pattern, var.separator)
  deduped_name = var.separator == "-" ? replace(local.normalized_name, "/-{2,}/", "-") : (
    var.separator == "_" ? replace(local.normalized_name, "/_{2,}/", "_") : local.normalized_name
  )
  full_name = var.separator == "" ? local.deduped_name : trim(local.deduped_name, var.separator)
  name      = var.separator == "" ? substr(local.full_name, 0, var.max_length) : trim(substr(local.full_name, 0, var.max_length), var.separator)

  labels_joined     = lower(join("-", local.parts))
  labels_normalized = replace(local.labels_joined, "/[^0-9a-z-]+/", "-")
  labels_deduped    = replace(local.labels_normalized, "/-{2,}/", "-")
  labels_trimmed    = trim(local.labels_deduped, "-")
  labels_safe_name  = trim(substr(local.labels_trimmed, 0, 63), "-")

  dns_joined     = lower(join("-", local.parts))
  dns_normalized = replace(local.dns_joined, "/[^0-9a-z-]+/", "-")
  dns_deduped    = replace(local.dns_normalized, "/-{2,}/", "-")
  dns_trimmed    = trim(local.dns_deduped, "-")
  dns_safe_name  = trim(substr(local.dns_trimmed, 0, 63), "-")
}
