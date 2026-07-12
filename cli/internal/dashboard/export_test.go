package dashboard

import (
	"encoding/json"
	cfapp "github.com/alekpopovic/clusterforge/cli/internal/app"
	"github.com/alekpopovic/clusterforge/cli/internal/config"
	"github.com/alekpopovic/clusterforge/cli/internal/servicecatalog"
	"strings"
	"testing"
)

func TestMinimalAndAppExportRedactsSecrets(t *testing.T) {
	cfg := config.DefaultConfig("demo")
	cfg.Environments["dev"] = config.Environment{Cloud: "aws", Region: "eu", Orchestrator: "eks", Path: "live/dev"}
	apps := map[string]cfapp.Manifest{"api": {Name: "api", Type: "web", Image: "example/api:1", Env: map[string]string{"PASSWORD": "do-not-export"}, SecretEnv: map[string]cfapp.SecretRef{"TOKEN": {SecretName: "private", SecretKey: "token"}}}}
	out := Build(cfg, apps, servicecatalog.Catalog{}, nil, "", "", nil)
	data, err := json.Marshal(out)
	if err != nil || !json.Valid(data) {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "do-not-export") || strings.Contains(string(data), "private") {
		t.Fatalf("secret leak: %s", data)
	}
	if out.SchemaVersion != "1.0" || len(out.Apps) != 1 {
		t.Fatalf("%#v", out)
	}
}
