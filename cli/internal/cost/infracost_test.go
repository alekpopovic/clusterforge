package cost

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestInfracostMissing(t *testing.T) {
	t.Setenv("PATH", t.TempDir())
	err := RunInfracost(context.Background(), ".", false, &bytes.Buffer{}, &bytes.Buffer{})
	if err == nil || !strings.Contains(err.Error(), "not installed") {
		t.Fatalf("%v", err)
	}
}
