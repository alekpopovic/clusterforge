package plugins

import (
	"bytes"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func writePlugin(t *testing.T, directory, name, body string) string {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("executable shell plugin fixture requires a POSIX shell")
	}
	path := filepath.Join(directory, prefix+name)
	if err := os.WriteFile(path, []byte("#!/bin/sh\n"+body+"\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestDiscoverFromConfiguredDirectory(t *testing.T) {
	directory := t.TempDir()
	writePlugin(t, directory, "hello", "exit 0")
	plugins, err := Discover(DiscoverOptions{Directories: []string{directory}})
	if err != nil {
		t.Fatal(err)
	}
	if len(plugins) != 1 || plugins[0].Name != "hello" {
		t.Fatalf("unexpected plugins: %#v", plugins)
	}
}

func TestNoPluginsDisablesDiscovery(t *testing.T) {
	directory := t.TempDir()
	writePlugin(t, directory, "hello", "exit 0")
	plugins, err := Discover(DiscoverOptions{Directories: []string{directory}, NoPlugins: true})
	if err != nil {
		t.Fatal(err)
	}
	if len(plugins) != 0 {
		t.Fatalf("expected no plugins, got %#v", plugins)
	}
}

func TestReadInfo(t *testing.T) {
	directory := t.TempDir()
	path := writePlugin(t, directory, "hello", `printf '%s\n' '{"name":"hello","version":"0.1.0","description":"test","commands":["greet"],"capabilities":["generator"]}'`)
	info, err := ReadInfo(Plugin{Name: "hello", Path: path})
	if err != nil {
		t.Fatal(err)
	}
	if info.Name != "hello" || info.Version != "0.1.0" || len(info.Capabilities) != 1 {
		t.Fatalf("unexpected info: %#v", info)
	}
}

func TestDisabledPluginIsNotExecuted(t *testing.T) {
	directory := t.TempDir()
	marker := filepath.Join(directory, "executed")
	path := writePlugin(t, directory, "hello", "touch "+marker)
	err := Run(Plugin{Name: "hello", Path: path, Disabled: true}, nil, nil, &bytes.Buffer{}, &bytes.Buffer{})
	if err == nil || !strings.Contains(err.Error(), "disabled") {
		t.Fatalf("expected disabled error, got %v", err)
	}
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Fatalf("disabled plugin executed: %v", err)
	}
}

func TestRunPassesArguments(t *testing.T) {
	directory := t.TempDir()
	path := writePlugin(t, directory, "hello", `printf '%s' "$*"`)
	var stdout bytes.Buffer
	if err := Run(Plugin{Name: "hello", Path: path}, []string{"one", "two"}, nil, &stdout, &bytes.Buffer{}); err != nil {
		t.Fatal(err)
	}
	if stdout.String() != "one two" {
		t.Fatalf("unexpected output %q", stdout.String())
	}
}
