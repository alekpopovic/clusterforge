package e2e

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var cfBinary string

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "clusterforge-e2e-")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)

	binary := filepath.Join(dir, "cf"+exeSuffix())
	cmd := exec.Command("go", "build", "-o", binary, "..")
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Fprintf(os.Stderr, "build cf: %v\n%s", err, output)
		os.Exit(1)
	}
	cfBinary = binary

	os.Exit(m.Run())
}

func TestProjectInitCreatesConfigAndDirectories(t *testing.T) {
	dir := t.TempDir()
	runCF(t, dir, "project", "init", "demo")

	assertExists(t, filepath.Join(dir, "clusterforge.yaml"))
	assertExists(t, filepath.Join(dir, "apps"))
	assertExists(t, filepath.Join(dir, "live"))
	assertExists(t, filepath.Join(dir, ".cf"))
}

func runCF(t *testing.T, dir string, args ...string) string {
	t.Helper()
	output, err := runCFAllowError(dir, args...)
	if err != nil {
		t.Fatalf("cf %s failed: %v\n%s", strings.Join(args, " "), err, output)
	}
	return output
}

func runCFAllowError(dir string, args ...string) (string, error) {
	allArgs := append([]string{"--non-interactive"}, args...)
	cmd := exec.Command(cfBinary, allArgs...)
	cmd.Dir = dir
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	err := cmd.Run()
	return output.String(), err
}

func assertExists(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected %s to exist: %v", path, err)
	}
}

func assertContains(t *testing.T, value, want string) {
	t.Helper()
	if !strings.Contains(value, want) {
		t.Fatalf("expected output to contain %q:\n%s", want, value)
	}
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return string(data)
}

func exeSuffix() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}
	return ""
}
