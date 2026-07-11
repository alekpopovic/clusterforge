package health

import (
	"context"
	"encoding/json"
	"testing"
)

func TestMissingPathFails(t *testing.T) {
	r := Evaluate(context.Background(), Input{Environment: "dev", Path: "/definitely/missing"})
	if r.Status != Fail {
		t.Fatalf("%#v", r)
	}
}
func TestLiveChecksSkipWithoutKubeconfig(t *testing.T) {
	r := Evaluate(context.Background(), Input{Environment: "dev", Path: t.TempDir(), CheckNodes: true, CheckAddons: true})
	if r.Status == Fail {
		t.Fatalf("%#v", r)
	}
	found := false
	for _, c := range r.Checks {
		if c.Status == Skipped {
			found = true
		}
	}
	if !found {
		t.Fatal("expected skipped live check")
	}
}
func TestJSONOutput(t *testing.T) {
	data, err := json.Marshal(Evaluate(context.Background(), Input{Environment: "dev", Path: t.TempDir()}))
	if err != nil || !json.Valid(data) {
		t.Fatalf("%s %v", data, err)
	}
}
