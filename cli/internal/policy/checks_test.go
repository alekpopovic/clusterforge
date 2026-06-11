package policy

import (
	"testing"

	"github.com/textracta/clusterforge/cli/internal/config"
)

func TestCheckApplyRequiresPlanFileForProd(t *testing.T) {
	err := CheckApply(Operation{
		Environment: "prod",
		ConfirmProd: true,
		Policies: config.Policies{
			RequirePlanFileForApply: true,
		},
	})
	if err == nil {
		t.Fatal("expected prod apply without plan file to fail")
	}
}

func TestCheckApplyAllowsDevWithoutPlanFile(t *testing.T) {
	err := CheckApply(Operation{
		Environment: "dev",
		Policies: config.Policies{
			RequirePlanFileForApply: true,
		},
	})
	if err != nil {
		t.Fatalf("expected dev apply to pass: %v", err)
	}
}

func TestCheckDestroyBlocksProdByDefault(t *testing.T) {
	err := CheckDestroy(Operation{
		Environment: "prod",
		Policies: config.Policies{
			BlockDestroyInProd: true,
		},
	})
	if err == nil {
		t.Fatal("expected prod destroy to be blocked")
	}
}

func TestCheckDestroyAllowsProdWithOverride(t *testing.T) {
	err := CheckDestroy(Operation{
		Environment:  "prod",
		AllowDestroy: true,
		ConfirmProd:  true,
		Policies: config.Policies{
			BlockDestroyInProd: true,
		},
	})
	if err != nil {
		t.Fatalf("expected prod destroy with override to pass: %v", err)
	}
}

func TestCheckApplyRequiresProdConfirmation(t *testing.T) {
	err := CheckApply(Operation{
		Environment: "prod",
		PlanFile:    "prod.tfplan",
		Policies: config.Policies{
			RequireManualApprovalForProd: true,
		},
	})
	if err == nil {
		t.Fatal("expected prod apply without confirmation to fail")
	}
}

func TestCheckDestroyRequiresProdConfirmation(t *testing.T) {
	err := CheckDestroy(Operation{
		Environment:  "prod",
		AllowDestroy: true,
		Policies: config.Policies{
			RequireManualApprovalForProd: true,
		},
	})
	if err == nil {
		t.Fatal("expected prod destroy without confirmation to fail")
	}
}
