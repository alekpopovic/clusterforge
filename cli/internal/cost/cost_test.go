package cost

import (
	"encoding/json"
	"testing"
)

func TestScanPlanJSONWarnings(t *testing.T) {
	data := []byte(`{
	  "resource_changes": [
	    {"address":"aws_nat_gateway.this","type":"aws_nat_gateway","change":{"actions":["create"]}},
	    {"address":"aws_eks_cluster.this","type":"aws_eks_cluster","change":{"actions":["create"]}},
	    {"address":"aws_vpc.this","type":"aws_vpc","change":{"actions":["create"]}}
	  ]
	}`)
	summary, err := ScanPlanJSON(data)
	if err != nil {
		t.Fatalf("ScanPlanJSON: %v", err)
	}
	if len(summary.Warnings) != 2 {
		t.Fatalf("warnings = %d, want 2", len(summary.Warnings))
	}
}

func TestScanPlanJSONRDSWarning(t *testing.T) {
	data := []byte(`{"resource_changes":[{"address":"aws_db_instance.db","type":"aws_db_instance","change":{"actions":["create"]}}]}`)
	summary, err := ScanPlanJSON(data)
	if err != nil {
		t.Fatal(err)
	}
	if len(summary.Warnings) != 1 || summary.Warnings[0].Category != "RDS database" {
		t.Fatalf("%#v", summary)
	}
}
func TestSummaryJSON(t *testing.T) {
	data, err := json.Marshal(Summary{Warnings: []Warning{{Category: "NAT Gateway"}}})
	if err != nil || !json.Valid(data) {
		t.Fatalf("%s %v", data, err)
	}
}

func TestScanPlanJSONEmpty(t *testing.T) {
	summary, err := ScanPlanJSON([]byte(`{"resource_changes":[]}`))
	if err != nil {
		t.Fatalf("ScanPlanJSON: %v", err)
	}
	if len(summary.Warnings) != 0 {
		t.Fatalf("warnings = %d, want 0", len(summary.Warnings))
	}
}
