package compliance

import (
	"bytes"
	"strings"
	"testing"
)

func TestReportRendersMarkdown(t *testing.T) {
	pack, err := Get("cis-kubernetes")
	if err != nil {
		t.Fatal(err)
	}
	var output bytes.Buffer
	if err := Render(&output, pack, "markdown"); err != nil {
		t.Fatal(err)
	}
	for _, expected := range []string{"# CIS Kubernetes", "| Control area |", "not certification"} {
		if !strings.Contains(output.String(), expected) {
			t.Fatalf("missing %q in %s", expected, output.String())
		}
	}
}

func TestUnknownPackFails(t *testing.T) {
	if _, err := Get("unknown"); err == nil {
		t.Fatal("expected unknown pack error")
	}
}
