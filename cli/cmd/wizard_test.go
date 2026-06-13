package cmd

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRequireValueNonInteractive(t *testing.T) {
	previous := opts.NonInteractive
	opts.NonInteractive = true
	t.Cleanup(func() {
		opts.NonInteractive = previous
	})

	_, err := requireValue(&cobra.Command{}, "", "project name")
	if err == nil {
		t.Fatal("expected missing value error")
	}
	if !strings.Contains(err.Error(), "project name is required in non-interactive mode") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPromptSessionReadsMultipleValues(t *testing.T) {
	input := strings.NewReader("aws\neks\nus-east-1\ny\n")
	output := &bytes.Buffer{}
	prompts := &promptSession{
		reader: bufio.NewReader(input),
		out:    output,
	}

	cloud, err := prompts.String("cloud", "aws")
	if err != nil {
		t.Fatalf("prompt cloud: %v", err)
	}
	orchestrator, err := prompts.String("orchestrator", "eks")
	if err != nil {
		t.Fatalf("prompt orchestrator: %v", err)
	}
	region, err := prompts.String("region", "eu-central-1")
	if err != nil {
		t.Fatalf("prompt region: %v", err)
	}
	enabled, err := prompts.Bool("enable autoscaling", false)
	if err != nil {
		t.Fatalf("prompt bool: %v", err)
	}

	if cloud != "aws" || orchestrator != "eks" || region != "us-east-1" || !enabled {
		t.Fatalf("prompts = %q %q %q %t", cloud, orchestrator, region, enabled)
	}
}
