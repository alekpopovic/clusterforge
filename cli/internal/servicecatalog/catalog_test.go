package servicecatalog

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func fixture() Catalog {
	return Catalog{Services: map[string]Service{"api": {Owner: "payments", Tier: "backend", Lifecycle: "production", Dependencies: []string{"postgres"}}, "postgres": {Owner: "platform", Lifecycle: "production"}}}
}
func TestValidateListExportAndGraph(t *testing.T) {
	c := fixture()
	if err := c.Validate(); err != nil {
		t.Fatal(err)
	}
	if c.Names()[0] != "api" {
		t.Fatal(c.Names())
	}
	data, err := c.JSON()
	if err != nil || !json.Valid(data) {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := c.WriteMarkdown(&out); err != nil || !strings.Contains(out.String(), "payments") {
		t.Fatal(out.String())
	}
	if !strings.Contains(c.DOT(), `"api" -> "postgres"`) {
		t.Fatal(c.DOT())
	}
}
