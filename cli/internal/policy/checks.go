package policy

import (
	"fmt"
	"strings"

	"github.com/textracta/clusterforge/cli/internal/config"
)

type Operation struct {
	Environment  string
	PlanFile     string
	AllowDestroy bool
	ConfirmProd  bool
	Policies     config.Policies
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
