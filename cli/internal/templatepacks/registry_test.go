package templatepacks

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseGitSource(t *testing.T) {
	parsed, err := ParseGitSource("git::https://example.com/templates.git?ref=v1.2.3")
	if err != nil {
		t.Fatal(err)
	}
	if parsed.URL != "https://example.com/templates.git" || parsed.Ref != "v1.2.3" {
		t.Fatalf("unexpected source: %#v", parsed)
	}
	if _, err := ParseGitSource("git::https://example.com/templates.git"); err == nil {
		t.Fatal("expected unpinned git source to fail")
	}
}

func TestWeakRef(t *testing.T) {
	if !WeakRef("main") || !WeakRef("master") || WeakRef("v1.2.3") {
		t.Fatal("unexpected weak ref classification")
	}
}

func TestCachePath(t *testing.T) {
	want := filepath.Join("repo", ".cf", "cache", "template-packs", "company", "v1")
	if got := CachePath("repo", "company", "v1"); got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestFetchLocalPath(t *testing.T) {
	source := filepath.Join(t.TempDir(), "source")
	if err := os.MkdirAll(filepath.Join(source, "env"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(source, "env", "main.tf.tmpl"), []byte("test"), 0o644); err != nil {
		t.Fatal(err)
	}
	destination := filepath.Join(t.TempDir(), "cache", "pack")
	if err := Fetch(source, destination); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(destination, "env", "main.tf.tmpl")); err != nil {
		t.Fatal(err)
	}
}
