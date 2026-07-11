package compliance

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

type Control struct {
	Area        string `json:"control_area"`
	Control     string `json:"control"`
	Feature     string `json:"clusterforge_feature_or_policy"`
	Status      string `json:"status"`
	Evidence    string `json:"evidence_source"`
	Limitations string `json:"limitations"`
}

type Pack struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Disclaimer string    `json:"disclaimer"`
	Controls   []Control `json:"controls"`
}

const disclaimer = "Control mapping and implementation aid only; not certification, attestation, or legal/compliance advice."

var packs = map[string]Pack{
	"soc2": {ID: "soc2", Name: "SOC 2 trust services criteria mapping", Disclaimer: disclaimer, Controls: []Control{
		{Area: "Logical access", Control: "CC6", Feature: "workload identity, RBAC, secret-reference policies", Status: "partial", Evidence: "Terraform config; policy result", Limitations: "Identity governance and access reviews remain organization-owned."},
		{Area: "System operations", Control: "CC7", Feature: "local audit log, drift and policy checks", Status: "implemented", Evidence: "audit log; CI result", Limitations: "Local logs are mutable and require external retention."},
		{Area: "Change management", Control: "CC8", Feature: "reviewable Terraform plans and production plan-file requirement", Status: "implemented", Evidence: "Terraform config; CI result", Limitations: "Approval workflow configuration is outside ClusterForge."},
	}},
	"cis-kubernetes": {ID: "cis-kubernetes", Name: "CIS Kubernetes Benchmark mapping", Disclaimer: disclaimer, Controls: []Control{
		{Area: "Control plane", Control: "API server and etcd settings", Feature: "managed-cluster module configuration", Status: "partial", Evidence: "Terraform config", Limitations: "Managed providers hide or own some benchmark settings."},
		{Area: "Workload security", Control: "Pod security standards", Feature: "Pod Security, Kyverno and Gatekeeper modules", Status: "implemented", Evidence: "Terraform config; policy result", Limitations: "Admission enforcement is opt-in and runtime coverage must be verified."},
		{Area: "Network policies", Control: "Namespace traffic isolation", Feature: "network-policy-baseline module", Status: "implemented", Evidence: "Terraform config; policy result", Limitations: "Application-specific allowed flows require manual design."},
	}},
	"aws-foundations": {ID: "aws-foundations", Name: "CIS AWS Foundations Benchmark mapping", Disclaimer: disclaimer, Controls: []Control{
		{Area: "Identity and access", Control: "IAM credential governance", Feature: "workload identity and no-static-secret conventions", Status: "partial", Evidence: "Terraform config; manual procedure", Limitations: "Organization-wide IAM, root account, and access reviews are not managed."},
		{Area: "Logging", Control: "CloudTrail and security monitoring", Feature: "AWS security policy pack guidance", Status: "documented", Evidence: "Terraform config; manual procedure", Limitations: "Account-wide trail deployment and alert response are environment-owned."},
		{Area: "Networking", Control: "Restrictive network access", Feature: "VPC, security group and private endpoint modules", Status: "partial", Evidence: "Terraform config; policy result", Limitations: "Benchmark evaluation requires account and region inventory."},
	}},
}

func List() []Pack {
	result := make([]Pack, 0, len(packs))
	for _, pack := range packs {
		result = append(result, pack)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].ID < result[j].ID })
	return result
}

func Get(id string) (Pack, error) {
	pack, ok := packs[id]
	if !ok {
		return Pack{}, fmt.Errorf("unknown compliance pack %q", id)
	}
	return pack, nil
}

func Render(writer io.Writer, pack Pack, format string) error {
	switch format {
	case "json":
		encoder := json.NewEncoder(writer)
		encoder.SetIndent("", "  ")
		return encoder.Encode(pack)
	case "markdown":
		fmt.Fprintf(writer, "# %s\n\n> %s\n\n", pack.Name, pack.Disclaimer)
		fmt.Fprintln(writer, "| Control area | Control | ClusterForge feature or policy | Status | Evidence source | Limitations |")
		fmt.Fprintln(writer, "|---|---|---|---|---|---|")
		for _, control := range pack.Controls {
			fmt.Fprintf(writer, "| %s | %s | %s | %s | %s | %s |\n", control.Area, control.Control, control.Feature, control.Status, control.Evidence, control.Limitations)
		}
		return nil
	default:
		return fmt.Errorf("unsupported format %q (use markdown or json)", format)
	}
}
