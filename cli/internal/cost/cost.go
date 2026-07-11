package cost

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Warning struct {
	Address  string `json:"address"`
	Type     string `json:"type"`
	Category string `json:"category"`
	Message  string `json:"message"`
}

type Summary struct {
	Warnings []Warning `json:"warnings"`
}

type plan struct {
	ResourceChanges []resourceChange `json:"resource_changes"`
}

type resourceChange struct {
	Address string `json:"address"`
	Type    string `json:"type"`
	Change  struct {
		Actions []string `json:"actions"`
	} `json:"change"`
}

func ScanPlanJSON(data []byte) (Summary, error) {
	var p plan
	if err := json.Unmarshal(data, &p); err != nil {
		return Summary{}, fmt.Errorf("parse plan JSON: %w", err)
	}
	var summary Summary
	for _, change := range p.ResourceChanges {
		if !hasCreateOrUpdate(change.Change.Actions) {
			continue
		}
		if warning, ok := warningFor(change); ok {
			summary.Warnings = append(summary.Warnings, warning)
		}
	}
	return summary, nil
}

func hasCreateOrUpdate(actions []string) bool {
	for _, action := range actions {
		switch action {
		case "create", "update", "replace":
			return true
		}
	}
	return false
}

func warningFor(change resourceChange) (Warning, bool) {
	resourceType := strings.ToLower(change.Type)
	categories := map[string]string{
		"aws_nat_gateway":                         "NAT Gateway",
		"aws_eks_cluster":                         "EKS control plane",
		"aws_eks_node_group":                      "EKS managed node group",
		"aws_lb":                                  "Load Balancer",
		"aws_alb":                                 "Load Balancer",
		"aws_elb":                                 "Load Balancer",
		"aws_ebs_volume":                          "Persistent volume",
		"kubernetes_persistent_volume_claim_v1":   "Persistent volume",
		"aws_cloudwatch_log_group":                "CloudWatch logs",
		"aws_db_instance":                         "RDS database",
		"aws_rds_cluster":                         "RDS database",
		"aws_elasticache_cluster":                 "ElastiCache",
		"aws_elasticache_replication_group":       "ElastiCache",
		"aws_s3_bucket_replication_configuration": "Cross-region replication",
		"aws_dynamodb_table_replica":              "Cross-region resource",
	}
	category, ok := categories[resourceType]
	if !ok {
		return Warning{}, false
	}
	return Warning{
		Address:  change.Address,
		Type:     change.Type,
		Category: category,
		Message:  fmt.Sprintf("%s can create ongoing cost; review sizing, retention, and cleanup", category),
	}, true
}
