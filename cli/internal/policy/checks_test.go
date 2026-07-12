package policy

import (
	"testing"

	"github.com/alekpopovic/clusterforge/cli/internal/config"
	"github.com/alekpopovic/clusterforge/cli/internal/terraform/planjson"
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

func TestCheckApplyPlanBlocksProdDeletes(t *testing.T) {
	_, err := CheckApplyPlan(Operation{
		Environment: "prod",
	}, planjson.Summary{Deletes: 1})
	if err == nil {
		t.Fatal("expected prod delete to be blocked")
	}
}

func TestCheckApplyPlanAllowsProdDeletesWithOverride(t *testing.T) {
	evaluation, err := CheckApplyPlan(Operation{
		Environment:  "prod",
		AllowDestroy: true,
	}, planjson.Summary{Deletes: 1})
	if err != nil {
		t.Fatalf("expected prod delete override to pass: %v", err)
	}
	if evaluation.Risk != RiskHigh {
		t.Fatalf("risk = %q", evaluation.Risk)
	}
}

func TestEvaluatePlanWarnsForProdReplacement(t *testing.T) {
	evaluation := EvaluatePlan("prod", planjson.Summary{Replacements: 1})
	if evaluation.Risk != RiskHigh {
		t.Fatalf("risk = %q", evaluation.Risk)
	}
	if len(evaluation.Warnings) == 0 {
		t.Fatalf("expected prod replacement warning")
	}
}

func TestEvaluatePlanReplacementInNonProdIsMedium(t *testing.T) {
	evaluation := EvaluatePlan("dev", planjson.Summary{Replacements: 1})
	if evaluation.Risk != RiskMedium {
		t.Fatalf("risk = %q", evaluation.Risk)
	}
}

func TestEvaluatePlanCreatesUpdatesInNonProdIsLow(t *testing.T) {
	evaluation := EvaluatePlan("dev", planjson.Summary{Creates: 1, Updates: 1})
	if evaluation.Risk != RiskLow {
		t.Fatalf("risk = %q", evaluation.Risk)
	}
}

func TestEvaluatePlanParseErrorFailsClosedForProd(t *testing.T) {
	evaluation, err := EvaluatePlanParseError("prod", errTestParse)
	if err == nil {
		t.Fatal("expected prod parse error to fail")
	}
	if evaluation.Risk != RiskBlocked {
		t.Fatalf("risk = %q", evaluation.Risk)
	}
}

func TestEvaluatePlanParseErrorWarnsForDev(t *testing.T) {
	evaluation, err := EvaluatePlanParseError("dev", errTestParse)
	if err != nil {
		t.Fatalf("expected dev parse error to warn only: %v", err)
	}
	if evaluation.Risk != RiskMedium || len(evaluation.Warnings) == 0 {
		t.Fatalf("evaluation = %#v", evaluation)
	}
}

var errTestParse = testError("bad json")

type testError string

func (e testError) Error() string {
	return string(e)
}
