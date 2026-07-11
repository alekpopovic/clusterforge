package migration

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAnalyzeTerraformFixture(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "live", "prod"), 0o755); err != nil {
		t.Fatal(err)
	}
	fixture := `terraform { backend "s3" {} }
provider "aws" {}
module "network" { source = "git::https://example.invalid/network.git" }
resource "aws_eks_cluster" "main" {}
resource "aws_ecs_cluster" "main" {}
database_password = "must-not-appear"
`
	if err := os.WriteFile(filepath.Join(root, "live", "prod", "main.tf"), []byte(fixture), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "live", "prod", "terraform.tfstate"), []byte("state-secret"), 0o600); err != nil {
		t.Fatal(err)
	}
	report, err := Analyze(root)
	if err != nil {
		t.Fatal(err)
	}
	if report.ResourceCounts["eks"] != 1 || report.ResourceCounts["ecs"] != 1 || len(report.Backends) != 1 || report.Backends[0] != "s3" {
		t.Fatalf("report=%#v", report)
	}
	data, _ := report.JSON()
	if strings.Contains(string(data), "must-not-appear") || !strings.Contains(string(data), "[REDACTED]") {
		t.Fatalf("secret redaction failed: %s", data)
	}
}
