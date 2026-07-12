package scaffold

import (
	"fmt"
	"strings"
)

type Plan struct {
	Project, Owner, Target, Environment, Backend           string
	Addons                                                 []string
	DemoApp                                                bool
	ProductionPlanRequired, DestroyBlocked, DisallowLatest bool
}

func Defaults() Plan {
	return Plan{Project: "clusterforge-demo", Owner: "platform-team", Target: "local-kind", Environment: "dev", Backend: "local", Addons: []string{"ingress", "cert-manager", "monitoring"}, DemoApp: true, ProductionPlanRequired: true, DestroyBlocked: true, DisallowLatest: true}
}
func (p Plan) Validate() error {
	if strings.TrimSpace(p.Project) == "" {
		return fmt.Errorf("project name is required")
	}
	if strings.TrimSpace(p.Owner) == "" {
		return fmt.Errorf("team/owner is required")
	}
	if !contains([]string{"aws-eks", "aws-ecs", "existing-kubernetes", "local-kind", "azure-aks", "gcp-gke"}, p.Target) {
		return fmt.Errorf("unsupported target %q", p.Target)
	}
	if !contains([]string{"dev", "staging", "prod"}, p.Environment) {
		return fmt.Errorf("environment must be dev, staging, or prod")
	}
	if !contains([]string{"local", "s3", "terraform-cloud"}, p.Backend) {
		return fmt.Errorf("backend must be local, s3, or terraform-cloud")
	}
	return nil
}
func (p Plan) Summary() string {
	return fmt.Sprintf("Project: %s\nOwner: %s\nTarget: %s\nEnvironment: %s\nBackend: %s\nAdd-ons: %s\nDemo app: %t\nSafety: production plan required=%t, destroy blocked=%t, disallow latest=%t", p.Project, p.Owner, p.Target, p.Environment, p.Backend, strings.Join(p.Addons, ", "), p.DemoApp, p.ProductionPlanRequired, p.DestroyBlocked, p.DisallowLatest)
}
func Target(target string) (cloud, orchestrator, region string) {
	switch target {
	case "aws-eks":
		return "aws", "eks", "us-east-1"
	case "aws-ecs":
		return "aws", "ecs", "us-east-1"
	case "azure-aks":
		return "azure", "aks", "westeurope"
	case "gcp-gke":
		return "gcp", "gke", "europe-west1"
	case "existing-kubernetes":
		return "existing", "kubernetes", "local"
	case "local-kind":
		// A kind cluster is consumed through the existing Kubernetes templates;
		// cluster lifecycle remains explicit through `cf local create kind`.
		return "existing", "kubernetes", "local"
	default:
		return "local", "kubernetes", "local"
	}
}
func contains(values []string, value string) bool {
	for _, candidate := range values {
		if candidate == value {
			return true
		}
	}
	return false
}
