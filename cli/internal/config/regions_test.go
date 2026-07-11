package config

import "testing"

func TestMultipleRegions(t *testing.T) {
	cfg := DefaultConfig("demo")
	cfg.Regions["primary"] = "eu-central-1"
	cfg.Regions["secondary"] = "eu-west-1"
	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
}
func TestDuplicateRegionAliasesFail(t *testing.T) {
	cfg := DefaultConfig("demo")
	cfg.Regions["a"] = "eu-central-1"
	cfg.Regions["b"] = "eu-central-1"
	if err := cfg.Validate(); err == nil {
		t.Fatal("expected duplicate region error")
	}
}
