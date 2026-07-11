package fleet

import (
	"context"
	"fmt"
	"strings"

	"github.com/textracta/clusterforge/cli/internal/inventory"
)

type Filter struct {
	Environment  string
	Cloud        string
	Orchestrator string
	Status       string
	Region       string
}

func Apply(clusters []inventory.Cluster, filter Filter) []inventory.Cluster {
	filtered := make([]inventory.Cluster, 0, len(clusters))
	for _, cluster := range clusters {
		if filter.Environment != "" && !strings.EqualFold(cluster.Environment, filter.Environment) {
			continue
		}
		if filter.Cloud != "" && !strings.EqualFold(cluster.Cloud, filter.Cloud) {
			continue
		}
		if filter.Orchestrator != "" && !strings.EqualFold(cluster.Orchestrator, filter.Orchestrator) {
			continue
		}
		if filter.Status != "" && !strings.EqualFold(cluster.Status, filter.Status) {
			continue
		}
		if filter.Region != "" && !strings.EqualFold(cluster.Region, filter.Region) {
			continue
		}
		filtered = append(filtered, cluster)
	}
	return filtered
}

type Result struct {
	Cluster   inventory.Cluster `json:"cluster"`
	Operation string            `json:"operation"`
	Status    string            `json:"status"`
	Message   string            `json:"message,omitempty"`
	Drift     bool              `json:"drift,omitempty"`
}

type Operation func(context.Context, inventory.Cluster) (Result, error)

func Run(ctx context.Context, clusters []inventory.Cluster, operationName string, failFast bool, operation Operation) ([]Result, error) {
	results := make([]Result, 0, len(clusters))
	for _, cluster := range clusters {
		result, err := operation(ctx, cluster)
		result.Cluster = cluster
		result.Operation = operationName
		if err != nil {
			result.Status = "fail"
			result.Message = err.Error()
			results = append(results, result)
			if failFast {
				return results, fmt.Errorf("%s failed for cluster %s: %w", operationName, cluster.Name, err)
			}
			continue
		}
		if result.Status == "" {
			result.Status = "pass"
		}
		results = append(results, result)
	}
	return results, nil
}
