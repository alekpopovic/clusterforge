locals {
  required_labels_template = <<-YAML
    apiVersion: templates.gatekeeper.sh/v1
    kind: ConstraintTemplate
    metadata:
      name: k8srequiredlabels
    spec:
      crd:
        spec:
          names:
            kind: K8sRequiredLabels
          validation:
            openAPIV3Schema:
              type: object
              properties:
                labels:
                  type: array
                  items:
                    type: string
      targets:
        - target: admission.k8s.gatekeeper.sh
          rego: |
            package k8srequiredlabels
            violation[{"msg": msg}] {
              provided := {label | input.review.object.metadata.labels[label]}
              required := {label | label := input.parameters.labels[_]}
              missing := required - provided
              count(missing) > 0
              msg := sprintf("missing required labels: %v", [missing])
            }
  YAML

  namespace_constraint = <<-YAML
    apiVersion: constraints.gatekeeper.sh/v1beta1
    kind: K8sRequiredLabels
    metadata:
      name: namespaces-require-owner
    spec:
      enforcementAction: dryrun
      match:
        kinds:
          - apiGroups: [""]
            kinds: ["Namespace"]
      parameters:
        labels: ["example.com/owner"]
  YAML
}

module "gatekeeper" {
  source = "../../modules/platform/kubernetes/gatekeeper"

  chart_version = var.gatekeeper_chart_version

  constraint_templates = var.enable_audit_constraint ? {
    required_labels = local.required_labels_template
  } : {}

  constraints = var.enable_audit_constraint ? {
    namespace_owner = local.namespace_constraint
  } : {}
}
