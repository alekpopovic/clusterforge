package config

import "testing"

func TestAWSAccountConfiguration(t *testing.T) {
	cfg := DefaultConfig("demo")
	cfg.AWSAccounts["prod"] = AWSAccount{AccountID: "222222222222", Region: "eu-central-1", Profile: "prod", RoleARN: "arn:aws:iam::222222222222:role/ClusterForgeDeployRole"}
	cfg.Environments["prod"] = Environment{Cloud: "aws", Region: "eu-central-1", Orchestrator: "eks", Path: "live/prod", Account: "prod"}
	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
}
