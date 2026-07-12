package backstage

import (
	cfapp "github.com/alekpopovic/clusterforge/cli/internal/app"
	"github.com/alekpopovic/clusterforge/cli/internal/config"
	"gopkg.in/yaml.v3"
	"strings"
	"testing"
)

func TestGenerateProjectAppAndOwnerFallback(t *testing.T) {
	cfg := config.DefaultConfig("payments")
	cfg.Backstage = config.Backstage{Enabled: true, Owner: "platform-team", System: "payments-platform"}
	cfg.Environments["prod"] = config.Environment{Cloud: "aws", Orchestrator: "eks", Path: "live/prod"}
	data, err := Generate(cfg, map[string]cfapp.Manifest{"api": {Name: "api"}}, "", "")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "platform-team") || !strings.Contains(string(data), "kind: Component") {
		t.Fatalf("%s", data)
	}
	var node yaml.Node
	if err := yaml.Unmarshal(data, &node); err != nil {
		t.Fatal(err)
	}
}
