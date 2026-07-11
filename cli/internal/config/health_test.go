package config

import "testing"

func TestHealthConfig(t *testing.T) {
	cfg := DefaultConfig("demo")
	cfg.Environments["prod"] = Environment{Cloud: "aws", Region: "x", Orchestrator: "eks", Path: "live/prod"}
	cfg.Health.Environments["prod"] = EnvironmentHealth{SLO: SLO{AvailabilityTarget: "99.9", LatencyTargetMS: 300, ErrorRateTarget: "1%"}, Checks: HealthChecks{KubernetesNodes: true}}
	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
}
