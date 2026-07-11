package upgradeplanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOneMinorAllowed(t *testing.T) {
	p := KubernetesUpgrade(KubernetesInput{Current: "1.29", Target: "1.30", Matrix: []byte("`1.30`")})
	if len(p.Blocking) != 0 {
		t.Fatalf("%#v", p)
	}
}
func TestMinorJumpBlocked(t *testing.T) {
	p := KubernetesUpgrade(KubernetesInput{Current: "1.29", Target: "1.31"})
	if p.Readiness != "blocked" {
		t.Fatalf("%#v", p)
	}
}
func TestUnsupportedTargetWarns(t *testing.T) {
	p := KubernetesUpgrade(KubernetesInput{Current: "1.29", Target: "1.30", Matrix: []byte("1.29")})
	if len(p.Warnings) == 0 {
		t.Fatal("expected warning")
	}
}
func TestDeprecatedAPIDetected(t *testing.T) {
	d := t.TempDir()
	os.WriteFile(filepath.Join(d, "old.yaml"), []byte("apiVersion: extensions/v1beta1"), 0644)
	p := KubernetesUpgrade(KubernetesInput{Current: "1.29", Target: "1.30", Root: d, Matrix: []byte("`1.30`")})
	if len(p.Warnings) == 0 {
		t.Fatal("expected deprecated API warning")
	}
}
