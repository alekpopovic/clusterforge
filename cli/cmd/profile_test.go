package cmd

import (
	"github.com/textracta/clusterforge/cli/internal/config"
	"testing"
)

func TestProdProfileRequiresPlanFile(t *testing.T) {
	if err := validateProfileApply("prod", config.ExecutionProfile{RequirePlanFile: true}, ""); err == nil {
		t.Fatal("expected plan-file error")
	}
}
