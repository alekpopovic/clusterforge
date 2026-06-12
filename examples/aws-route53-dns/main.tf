module "tags" {
  source = "../../modules/core/tags"

  project     = "clusterforge"
  environment = "dev"
  component   = "dns"
}

module "dns" {
  source = "../../modules/cloud/aws/dns"

  create_zone = false
  zone_name   = "example.com"
  zone_id     = var.zone_id

  records = {
    txt = {
      name    = "clusterforge.example.com"
      type    = "TXT"
      ttl     = 300
      records = ["\"managed-by=clusterforge\""]
    }

    app = {
      name = "app.example.com"
      type = "A"

      alias = {
        name    = var.alb_dns_name
        zone_id = var.alb_zone_id
      }
    }
  }

  tags = module.tags.tags
}
