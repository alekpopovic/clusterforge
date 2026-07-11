package config

import "testing"

func bp(v bool) *bool { return &v }
func TestExecutionProfileParsing(t *testing.T) {
	cfg := DefaultConfig("demo")
	cfg.ExecutionProfiles["prod"] = ExecutionProfile{Engine: "terraform", Parallelism: 3, Refresh: bp(true), Input: bp(false), LockTimeout: "20m", RequirePlanFile: true}
	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
}
