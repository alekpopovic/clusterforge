package upgradeplanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPlatformVersionDiffAndCRD(t *testing.T) {
	d := t.TempDir()
	os.WriteFile(filepath.Join(d, "main.tf"), []byte(`resource "helm_release" "cert_manager" { name="cert-manager" version="v1.13.0" }`), 0644)
	p := PlatformUpgrade(d, map[string]string{"cert-manager": "v1.14.0"})
	if len(p.Changes) != 1 || !p.Changes[0].Change || !p.Changes[0].CRD || p.Changes[0].Warning == "" {
		t.Fatalf("%#v", p)
	}
}
func TestUnknownComponentWarns(t *testing.T) {
	p := PlatformUpgrade(t.TempDir(), map[string]string{"mystery": "1.0"})
	if len(p.Warnings) == 0 {
		t.Fatal("expected warning")
	}
}
