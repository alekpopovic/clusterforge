package cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alekpopovic/clusterforge/cli/internal/config"
)

func TestWorkspaceAndTeamList(t *testing.T) {
	original := opts
	t.Cleanup(func() { opts = original })
	opts.ConfigPath = filepath.Join(t.TempDir(), "clusterforge.yaml")
	cfg := config.DefaultConfig("demo")
	cfg.Environments["dev"] = config.Environment{Cloud: "aws", Region: "x", Orchestrator: "eks", Path: "live/dev"}
	cfg.Workspaces["platform"] = config.Workspace{Description: "Core", Environments: []string{"dev"}}
	cfg.Teams["platform"] = config.Team{Owners: []string{"platform@example.com"}}
	if err := cfg.Save(opts.ConfigPath); err != nil {
		t.Fatal(err)
	}
	var workspaceOut, teamOut bytes.Buffer
	workspaceListCmd.SetOut(&workspaceOut)
	teamListCmd.SetOut(&teamOut)
	t.Cleanup(func() { workspaceListCmd.SetOut(nil); teamListCmd.SetOut(nil) })
	if err := workspaceListCmd.RunE(workspaceListCmd, nil); err != nil {
		t.Fatal(err)
	}
	if err := teamListCmd.RunE(teamListCmd, nil); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(workspaceOut.String(), "platform") || !strings.Contains(teamOut.String(), "platform") {
		t.Fatalf("unexpected output: %q %q", workspaceOut.String(), teamOut.String())
	}
}
