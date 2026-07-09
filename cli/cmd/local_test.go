package cmd

import "testing"

func TestLocalTargetSpecKind(t *testing.T) {
	spec, err := localTargetSpec("kind")
	if err != nil {
		t.Fatalf("localTargetSpec: %v", err)
	}
	if spec.Binary != "kind" {
		t.Fatalf("binary = %q", spec.Binary)
	}
	want := []string{"kind", "create", "cluster", "--name", localClusterName}
	if len(spec.CreateArgs) != len(want) {
		t.Fatalf("args = %#v", spec.CreateArgs)
	}
	for i := range want {
		if spec.CreateArgs[i] != want[i] {
			t.Fatalf("arg %d = %q, want %q", i, spec.CreateArgs[i], want[i])
		}
	}
}

func TestLocalTargetSpecK3d(t *testing.T) {
	spec, err := localTargetSpec("k3d")
	if err != nil {
		t.Fatalf("localTargetSpec: %v", err)
	}
	if spec.Binary != "k3d" {
		t.Fatalf("binary = %q", spec.Binary)
	}
	if spec.DeleteArgs[0] != "k3d" || spec.DeleteArgs[1] != "cluster" || spec.DeleteArgs[2] != "delete" {
		t.Fatalf("delete args = %#v", spec.DeleteArgs)
	}
}

func TestLocalTargetSpecInvalid(t *testing.T) {
	if _, err := localTargetSpec("minikube"); err == nil {
		t.Fatal("expected unsupported target to fail")
	}
}
