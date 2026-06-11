package terraform

import (
	"reflect"
	"testing"
)

func TestPlanArgs(t *testing.T) {
	got := planArgs("dev.tfplan", []string{"-input=false"})
	want := []string{"plan", "-out", "dev.tfplan", "-input=false"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("planArgs() = %#v, want %#v", got, want)
	}
}

func TestPlanArgsWithoutOutFile(t *testing.T) {
	got := planArgs("", []string{"-refresh=false"})
	want := []string{"plan", "-refresh=false"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("planArgs() = %#v, want %#v", got, want)
	}
}

func TestApplyArgsWithPlanFile(t *testing.T) {
	got := applyArgs(".cf/plans/dev.tfplan", []string{"-input=false"})
	want := []string{"apply", ".cf/plans/dev.tfplan", "-input=false"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("applyArgs() = %#v, want %#v", got, want)
	}
}

func TestApplyArgsWithoutAutoApprove(t *testing.T) {
	got := applyArgs("", nil)
	want := []string{"apply"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("applyArgs() = %#v, want %#v", got, want)
	}
}

func TestDestroyArgsWithoutAutoApprove(t *testing.T) {
	got := destroyArgs(nil)
	want := []string{"destroy"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("destroyArgs() = %#v, want %#v", got, want)
	}
}

func TestOutputArgsJSON(t *testing.T) {
	got := outputArgs(true)
	want := []string{"output", "-json"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("outputArgs() = %#v, want %#v", got, want)
	}
}

func TestShowPlanJSONArgs(t *testing.T) {
	got := showPlanJSONArgs("dev.tfplan")
	want := []string{"show", "-json", "dev.tfplan"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("showPlanJSONArgs() = %#v, want %#v", got, want)
	}
}
