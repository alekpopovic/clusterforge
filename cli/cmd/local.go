package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

const localClusterName = "clusterforge-local"

var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Manage local Kubernetes development clusters",
}

var localCreateCmd = &cobra.Command{
	Use:   "create <kind|k3d>",
	Short: "Create a local Kubernetes cluster and generated environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		target := args[0]
		spec, err := localTargetSpec(target)
		if err != nil {
			return err
		}
		if _, err := exec.LookPath(spec.Binary); err != nil {
			return fmt.Errorf("%s binary is required for local %s; install it and retry", spec.Binary, target)
		}
		if err := runLocalCommand(cmd.Context(), spec.CreateArgs); err != nil {
			return err
		}
		if err := writeLocalEnvironment(target); err != nil {
			return err
		}
		printer.Success(fmt.Sprintf("created local %s cluster and live/local/%s", target, target))
		return nil
	},
}

var localDeleteCmd = &cobra.Command{
	Use:   "delete <kind|k3d>",
	Short: "Delete a named local Kubernetes cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		spec, err := localTargetSpec(args[0])
		if err != nil {
			return err
		}
		if _, err := exec.LookPath(spec.Binary); err != nil {
			return fmt.Errorf("%s binary is required for local delete", spec.Binary)
		}
		return runLocalCommand(cmd.Context(), spec.DeleteArgs)
	},
}

var localKubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig <kind|k3d>",
	Short: "Print kubeconfig command for a local cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		spec, err := localTargetSpec(args[0])
		if err != nil {
			return err
		}
		fmt.Fprintln(cmd.OutOrStdout(), spec.KubeconfigCommand)
		return nil
	},
}

var localStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show local Kubernetes cluster status",
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, target := range []string{"kind", "k3d"} {
			spec, _ := localTargetSpec(target)
			if _, err := exec.LookPath(spec.Binary); err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "%s: %s not installed\n", target, spec.Binary)
				continue
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%s: %s installed\n", target, spec.Binary)
		}
		return nil
	},
}

type localSpec struct {
	Binary            string
	CreateArgs        []string
	DeleteArgs        []string
	KubeconfigCommand string
}

func localTargetSpec(target string) (localSpec, error) {
	switch target {
	case "kind":
		return localSpec{
			Binary:            "kind",
			CreateArgs:        []string{"kind", "create", "cluster", "--name", localClusterName},
			DeleteArgs:        []string{"kind", "delete", "cluster", "--name", localClusterName},
			KubeconfigCommand: "kind get kubeconfig --name " + localClusterName,
		}, nil
	case "k3d":
		return localSpec{
			Binary:            "k3d",
			CreateArgs:        []string{"k3d", "cluster", "create", localClusterName},
			DeleteArgs:        []string{"k3d", "cluster", "delete", localClusterName},
			KubeconfigCommand: "k3d kubeconfig get " + localClusterName,
		}, nil
	default:
		return localSpec{}, fmt.Errorf("unsupported local target %q; expected kind or k3d", target)
	}
}

func runLocalCommand(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("local command args are required")
	}
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s failed: %w", args[0], err)
	}
	return nil
}

func writeLocalEnvironment(target string) error {
	dir := filepath.Join("live", "local", target)
	files := map[string]string{
		"versions.tf": `terraform {
  required_version = ">= 1.6.0"

  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 2.20, < 3.0"
    }
  }
}
`,
		"providers.tf": `provider "kubernetes" {
  config_path    = var.kubeconfig_path
  config_context = var.kubeconfig_context
}
`,
		"main.tf": `# Local Kubernetes environment for ClusterForge workloads.
# Render applications into this directory with:
#   cf app render <name> --env local
`,
		"variables.tf": `variable "kubeconfig_path" {
  description = "Path to kubeconfig for the local cluster."
  type        = string
  default     = "~/.kube/config"
}

variable "kubeconfig_context" {
  description = "Kubeconfig context for the local cluster."
  type        = string
  default     = ""
}
`,
		"outputs.tf": `output "kubeconfig_context" {
  description = "Configured kubeconfig context."
  value       = var.kubeconfig_context
}
`,
		"terraform.tfvars.example": fmt.Sprintf(`kubeconfig_path    = "~/.kube/config"
kubeconfig_context = "%s-%s"
`, target, localClusterName),
		"README.md": fmt.Sprintf(`# Local %s

Generated local Kubernetes environment for ClusterForge development.

This root uses the current kubeconfig and creates no cloud resources.
`, target),
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create local environment directory: %w", err)
	}
	for name, content := range files {
		path := filepath.Join(dir, name)
		if _, err := os.Stat(path); err == nil {
			continue
		}
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}
	}
	return nil
}

func init() {
	localCmd.AddCommand(localCreateCmd)
	localCmd.AddCommand(localDeleteCmd)
	localCmd.AddCommand(localKubeconfigCmd)
	localCmd.AddCommand(localStatusCmd)
}
