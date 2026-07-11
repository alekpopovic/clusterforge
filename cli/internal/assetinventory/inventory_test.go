package assetinventory

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

var stateFixture = []byte(`{"values":{"root_module":{"resources":[{"address":"aws_db_instance.db","mode":"managed","type":"aws_db_instance","name":"db","provider_name":"registry.terraform.io/hashicorp/aws","values":{"tags":{"Name":"db","secret-token":"should-not-print"},"password":"hidden"},"sensitive_values":{"password":true}}]}}}`)

func TestStateFixtureRedactsSensitive(t *testing.T) {
	assets, err := ParseState(stateFixture, "prod", "data", "aws", "eu-central-1")
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(assets)
	if strings.Contains(string(data), "should-not-print") || strings.Contains(string(data), "hidden") {
		t.Fatalf("leak: %s", data)
	}
	if !assets[0].Sensitive {
		t.Fatal("expected sensitive")
	}
}
func TestCSVAndJSON(t *testing.T) {
	assets, _ := ParseState(stateFixture, "prod", "", "aws", "eu")
	var out bytes.Buffer
	if err := WriteCSV(&out, assets); err != nil || !strings.Contains(out.String(), "aws_db_instance") {
		t.Fatal(out.String())
	}
	if data, err := json.Marshal(assets); err != nil || !json.Valid(data) {
		t.Fatal(err)
	}
}
