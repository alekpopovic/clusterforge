module "labels" {
  source = "../../modules/core/labels"

  app       = "hello"
  component = "web"
}

module "app" {
  source = "../../modules/workloads/kubernetes/app"

  name   = "hello"
  image  = "nginx:1.27"
  labels = module.labels.labels
}
