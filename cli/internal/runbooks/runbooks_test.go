package runbooks

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func fixture(t *testing.T) string {
	d := t.TempDir()
	data := []byte("---\ntitle: Failed Apply\ntags: [terraform, incident]\n---\n# Fallback\n\nActionable summary.\n")
	if err := os.WriteFile(filepath.Join(d, "failed.md"), data, 0644); err != nil {
		t.Fatal(err)
	}
	return d
}
func TestListSearchShow(t *testing.T) {
	all, err := Discover(fixture(t))
	if err != nil || len(all) != 1 {
		t.Fatalf("%#v %v", all, err)
	}
	if len(Search(all, "terraform")) != 1 {
		t.Fatal("search failed")
	}
	book, err := Find(all, "failed")
	if err != nil || !strings.Contains(book.Content, "Actionable") {
		t.Fatal("show failed")
	}
}
func TestMissingRunbook(t *testing.T) {
	if _, err := Find(nil, "missing"); err == nil {
		t.Fatal("expected error")
	}
}
