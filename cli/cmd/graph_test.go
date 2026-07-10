package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteGraphOutputFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "graphs", "dev.dot")
	if err := writeGraphOutput(&bytes.Buffer{}, path, []byte("digraph dev {}\n")); err != nil {
		t.Fatalf("write graph: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "digraph dev {}\n" {
		t.Fatalf("data = %q", data)
	}
}
