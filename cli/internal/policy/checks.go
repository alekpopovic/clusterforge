package policy

import (
	"fmt"
	"strings"

	"github.com/textracta/clusterforge/cli/internal/config"
	"github.com/textracta/clusterforge/cli/internal/terraform/planjson"
)

const (
	RiskLow     = "LOW"
	RiskMedium  = "MEDIUM"
	RiskHigh    = "HIGH"
	RiskBlocked = "BLOCKED"
)

type Operation struct {
	Environment  string
	PlanFile     string
	AllowDestroy bool
	ConfirmProd  bool
	Policies     config.Policies
}

type Evaluation struct {
	Risk     string
	Policy   string
	Warnings []string
	Blocked  bool
}

func CheckApply(op Operation) error {
	if isProd(op.Environment) && op.Policies.RequirePlanFileForApply && op.PlanFile == "" {
		return fmt.Errorf("apply against prod requires --plan-file")
	}
	if isProd(op.Environment) && op.Policies.RequireManualApprovalForProd && !op.ConfirmProd {
		return fmt.Errorf("apply against prod requires --confirm-prod")
	}
	return nil
}

func CheckApplyPlan(op Operation, summary planjson.Summary) (Evaluation, error) {
	evaluation := EvaluatePlan(op.Environment, summary)
	if isProd(op.Environment) && summary.Deletes > 0 && !op.AllowDestroy {
		evaluation.Risk = RiskBlocked
		evaluation.Blocked = true
		evaluation.Policy = "delete actions in prod require --allow-destroy"
		return evaluation, fmt.Errorf("plan contains delete actions in prod; pass --allow-destroy to continue")
	}
	return evaluation, nil
}

func EvaluatePlan(environment string, summary planjson.Summary) Evaluation {
	switch {
	case isProd(environment) && (summary.Deletes > 0 || summary.Replacements > 0):
		evaluation := Evaluation{Risk: RiskHigh, Policy: "prod delete/replacement changes require careful review"}
		if summary.Replacements > 0 {
			evaluation.Warnings = append(evaluation.Warnings, "replacement actions in prod are high risk")
		}
		return evaluation
	case isProd(environment) && (summary.Creates > 0 || summary.Updates > 0):
		return Evaluation{Risk: RiskMedium, Policy: "apply allowed only with plan file"}
	case !isProd(environment) && summary.Replacements > 0:
		return Evaluation{Risk: RiskMedium, Policy: "replacement changes require review"}
	default:
		return Evaluation{Risk: RiskLow, Policy: "apply allowed"}
	}
}

func EvaluatePlanParseError(environment string, err error) (Evaluation, error) {
	if isProd(environment) {
		return Evaluation{
			Risk:    RiskBlocked,
			Policy:  "failed closed because prod plan JSON could not be parsed",
			Blocked: true,
		}, fmt.Errorf("cannot parse plan JSON for prod: %w", err)
	}
	return Evaluation{
		Risk:     RiskMedium,
		Policy:   "plan JSON could not be parsed; continue only after manual review",
		Warnings: []string{fmt.Sprintf("could not parse plan JSON: %v", err)},
	}, nil
}

func CheckDestroy(op Operation) error {
	if isProd(op.Environment) && op.Policies.BlockDestroyInProd && !op.AllowDestroy {
		return fmt.Errorf("destroy against prod is blocked by default; pass --allow-destroy to continue")
	}
	if isProd(op.Environment) && op.Policies.RequireManualApprovalForProd && !op.ConfirmProd {
		return fmt.Errorf("destroy against prod requires --confirm-prod")
	}
	return nil
}

func isProd(environment string) bool {
	env := strings.ToLower(environment)
	return env == "prod" || env == "production"
}

func IsProd(environment string) bool {
	return isProd(environment)
}
