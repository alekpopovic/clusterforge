provider "kubernetes" {}
provider "helm" {
  kubernetes {}
}

module "gatekeeper" {
  source = "../../modules/platform/kubernetes/gatekeeper"

  chart_version      = var.gatekeeper_chart_version
  enforcement_action = var.enforce ? "deny" : "dryrun"
  constraint_templates = {
    required_labels = file("${path.module}/required-labels-template.yaml")
  }
  constraints = {
    public_ingress_approval = file("${path.module}/public-ingress-approval.yaml")
  }
}
