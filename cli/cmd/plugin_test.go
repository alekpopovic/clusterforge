package cmd

import (
	"io"
	"path/filepath"
	"testing"

	"github.com/textracta/clusterforge/cli/internal/config"
)

func TestNoPluginsGlobalOptionDisablesDiscovery(t *testing.T) {
	original := opts
	t.Cleanup(func() { opts = original })
	opts.NoPlugins = true
	opts.ConfigPath = filepath.Join(t.TempDir(), "missing.yaml")
	found, err := discoverPlugins()
	if err != nil {
		t.Fatal(err)
	}
	if len(found) != 0 {
		t.Fatalf("expected no plugins, got %#v", found)
	}
}

func TestCIRequiresExplicitPluginAllowance(t *testing.T) {
	original := opts
	t.Cleanup(func() { opts = original })
	t.Setenv("CI", "true")
	opts.NoPlugins = false
	opts.AllowPlugins = false
	if _, err := discoverPlugins(); err == nil {
		t.Fatal("expected CI plugin restriction")
	}
}

func TestPluginDisablePersistsConfiguration(t *testing.T) {
	original := opts
	t.Cleanup(func() { opts = original })
	directory := t.TempDir()
	opts.ConfigPath = filepath.Join(directory, "clusterforge.yaml")
	cfg := config.DefaultConfig("test")
	cfg.Plugins.Enabled = true
	if err := cfg.Save(opts.ConfigPath); err != nil {
		t.Fatal(err)
	}
	command := pluginToggleCommand(false)
	command.SetArgs([]string{"hello"})
	command.SetOut(io.Discard)
	if err := command.Execute(); err != nil {
		t.Fatal(err)
	}
	loaded, err := config.Load(opts.ConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	if !containsString(loaded.Plugins.Disabled, "hello") {
		t.Fatalf("plugin was not disabled: %#v", loaded.Plugins)
	}
}
