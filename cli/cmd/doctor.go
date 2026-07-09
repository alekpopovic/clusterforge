package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/textracta/clusterforge/cli/internal/config"
	"github.com/textracta/clusterforge/cli/internal/policy"
	"github.com/textracta/clusterforge/cli/internal/ui"
)

const terraformMinimumVersion = "1.6.0"

var doctorJSON bool

type doctorStatus string

const (
	doctorPass doctorStatus = "pass"
	doctorWarn doctorStatus = "warn"
	doctorFail doctorStatus = "fail"
)

type doctorCheck struct {
	Name    string       `json:"name"`
	Status  doctorStatus `json:"status"`
	Message string       `json:"message"`
}

type doctorReport struct {
	Version string        `json:"version"`
	Commit  string        `json:"commit"`
	Date    string        `json:"date"`
	Status  doctorStatus  `json:"status"`
	Checks  []doctorCheck `json:"checks"`
}

type doctorCommandRunner interface {
	LookPath(file string) (string, error)
	Output(ctx context.Context, name string, args ...string) ([]byte, error)
}

type realDoctorRunner struct{}

func (realDoctorRunner) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

func (realDoctorRunner) Output(ctx context.Context, name string, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	return cmd.CombinedOutput()
}

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Diagnose local ClusterForge prerequisites and project health",
	RunE: func(cmd *cobra.Command, args []string) error {
		report := runDoctor(cmd.Context(), opts.ConfigPath, realDoctorRunner{})
		if doctorJSON {
			if err := ui.WriteJSON(cmd.OutOrStdout(), report); err != nil {
				return err
			}
		} else {
			printDoctorReport(cmd.OutOrStdout(), report)
		}
		if doctorHasFailure(report) {
			return fmt.Errorf("doctor found one or more hard failures")
		}
		return nil
	},
}

func init() {
	doctorCmd.Flags().BoolVar(&doctorJSON, "json", false, "Print doctor results as JSON")
}

func runDoctor(ctx context.Context, configPath string, runner doctorCommandRunner) doctorReport {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	report := doctorReport{
		Version: Version,
		Commit:  Commit,
		Date:    Date,
	}

	var cfg *config.Config
	report.Checks = append(report.Checks, checkCLIVersion()...)
	report.Checks = append(report.Checks, checkBinaries(ctx, runner)...)
	report.Checks = append(report.Checks, checkTerraform(ctx, runner)...)
	report.Checks = append(report.Checks, checkKubernetesVersion(ctx, runner)...)
	cfgChecks, loadedConfig := checkProjectConfig(configPath)
	report.Checks = append(report.Checks, cfgChecks...)
	cfg = loadedConfig
	if cfg != nil {
		report.Checks = append(report.Checks, checkProjectPaths(cfg)...)
		report.Checks = append(report.Checks, checkSafety(cfg)...)
	}
	report.Checks = append(report.Checks, checkGit(ctx, runner)...)
	report.Status = reportStatus(report.Checks)
	return report
}

func checkCLIVersion() []doctorCheck {
	return []doctorCheck{
		{
			Name:    "cli.version",
			Status:  doctorPass,
			Message: fmt.Sprintf("version=%s commit=%s date=%s", Version, Commit, Date),
		},
	}
}

func checkBinaries(ctx context.Context, runner doctorCommandRunner) []doctorCheck {
	checks := []doctorCheck{
		checkBinary(runner, "binary.terraform", "terraform", true),
		checkBinary(runner, "binary.tofu", "tofu", false),
		checkBinary(runner, "binary.git", "git", true),
		checkBinary(runner, "binary.kubectl", "kubectl", false),
		checkBinary(runner, "binary.helm", "helm", false),
		checkBinary(runner, "binary.go", "go", false),
	}
	if _, err := runner.Output(ctx, "go", "version"); err == nil {
		checks = append(checks, doctorCheck{Name: "cli.development", Status: doctorPass, Message: "go can run; CLI development checks available"})
	}
	return checks
}

func checkBinary(runner doctorCommandRunner, name, binary string, required bool) doctorCheck {
	path, err := runner.LookPath(binary)
	if err != nil {
		status := doctorWarn
		message := fmt.Sprintf("optional binary %q not found", binary)
		if required {
			status = doctorFail
			message = fmt.Sprintf("required binary %q not found", binary)
		}
		return doctorCheck{Name: name, Status: status, Message: message}
	}
	return doctorCheck{Name: name, Status: doctorPass, Message: fmt.Sprintf("%s found at %s", binary, path)}
}

func checkTerraform(ctx context.Context, runner doctorCommandRunner) []doctorCheck {
	output, err := runner.Output(ctx, "terraform", "version")
	if err != nil {
		return []doctorCheck{{
			Name:    "terraform.version",
			Status:  doctorFail,
			Message: fmt.Sprintf("terraform version failed: %v", err),
		}}
	}
	version, ok := parseTerraformVersion(string(output))
	if !ok {
		return []doctorCheck{{
			Name:    "terraform.version",
			Status:  doctorWarn,
			Message: "could not parse terraform version output",
		}}
	}
	status := doctorPass
	message := fmt.Sprintf("terraform %s", version)
	if compareSemver(version, terraformMinimumVersion) < 0 {
		status = doctorWarn
		message = fmt.Sprintf("terraform %s is below recommended minimum %s", version, terraformMinimumVersion)
	}
	return []doctorCheck{{Name: "terraform.version", Status: status, Message: message}}
}

func checkKubernetesVersion(ctx context.Context, runner doctorCommandRunner) []doctorCheck {
	if _, err := runner.LookPath("kubectl"); err != nil {
		return nil
	}
	output, err := runner.Output(ctx, "kubectl", "version", "--client=true")
	if err != nil {
		return []doctorCheck{{Name: "kubernetes.version", Status: doctorWarn, Message: fmt.Sprintf("kubectl version failed: %v", err)}}
	}
	version, ok := parseKubernetesVersion(string(output))
	if !ok {
		return []doctorCheck{{Name: "kubernetes.version", Status: doctorWarn, Message: "could not parse kubectl client version"}}
	}
	minor := semverParts(version)[1]
	if minor < 29 || minor > 31 {
		return []doctorCheck{{Name: "kubernetes.version", Status: doctorWarn, Message: fmt.Sprintf("kubectl client %s is outside the documented 1.29-1.31 tested matrix", version)}}
	}
	return []doctorCheck{{Name: "kubernetes.version", Status: doctorPass, Message: fmt.Sprintf("kubectl client %s is inside the documented tested matrix", version)}}
}

func checkProjectConfig(configPath string) ([]doctorCheck, *config.Config) {
	checks := []doctorCheck{}
	if _, err := os.Stat(configPath); err != nil {
		return []doctorCheck{{
			Name:    "project.config_file",
			Status:  doctorFail,
			Message: fmt.Sprintf("%s does not exist", configPath),
		}}, nil
	}
	checks = append(checks, doctorCheck{Name: "project.config_file", Status: doctorPass, Message: fmt.Sprintf("%s exists", configPath)})

	cfg, err := config.Load(configPath)
	if err != nil {
		checks = append(checks, doctorCheck{Name: "project.config_load", Status: doctorFail, Message: err.Error()})
		return checks, nil
	}
	checks = append(checks, doctorCheck{Name: "project.config_load", Status: doctorPass, Message: "config loaded"})
	if strings.TrimSpace(cfg.Project.Name) == "" {
		checks = append(checks, doctorCheck{Name: "project.name", Status: doctorFail, Message: "project.name is empty"})
	} else {
		checks = append(checks, doctorCheck{Name: "project.name", Status: doctorPass, Message: cfg.Project.Name})
	}
	if len(cfg.Environments) == 0 {
		checks = append(checks, doctorCheck{Name: "project.environments", Status: doctorFail, Message: "no environments configured"})
	} else {
		checks = append(checks, doctorCheck{Name: "project.environments", Status: doctorPass, Message: fmt.Sprintf("%d environment(s) configured", len(cfg.Environments))})
	}
	return checks, cfg
}

func checkProjectPaths(cfg *config.Config) []doctorCheck {
	checks := []doctorCheck{}
	names := sortedEnvironmentNames(cfg)
	for _, name := range names {
		env := cfg.Environments[name]
		if _, err := os.Stat(env.Path); err != nil {
			checks = append(checks, doctorCheck{
				Name:    fmt.Sprintf("environment.%s.path", name),
				Status:  doctorFail,
				Message: fmt.Sprintf("%s does not exist", env.Path),
			})
			continue
		}
		checks = append(checks, doctorCheck{
			Name:    fmt.Sprintf("environment.%s.path", name),
			Status:  doctorPass,
			Message: fmt.Sprintf("%s exists", env.Path),
		})
	}
	return checks
}

func checkSafety(cfg *config.Config) []doctorCheck {
	checks := []doctorCheck{}
	hasProd := false
	for _, name := range sortedEnvironmentNames(cfg) {
		if !policy.IsProd(name) {
			continue
		}
		hasProd = true
		backend := cfg.BackendFor(name)
		if backend.EffectiveType() == "local" {
			checks = append(checks, doctorCheck{
				Name:    fmt.Sprintf("safety.%s.backend", name),
				Status:  doctorWarn,
				Message: "prod environment uses local backend; configure a remote backend with locking",
			})
		} else {
			checks = append(checks, doctorCheck{
				Name:    fmt.Sprintf("safety.%s.backend", name),
				Status:  doctorPass,
				Message: fmt.Sprintf("prod backend type is %s", backend.EffectiveType()),
			})
		}
	}
	if !hasProd {
		checks = append(checks, doctorCheck{Name: "safety.prod_backend", Status: doctorPass, Message: "no prod environment configured"})
	}
	if cfg.Policies.BlockDestroyInProd {
		checks = append(checks, doctorCheck{Name: "safety.block_destroy_in_prod", Status: doctorPass, Message: "enabled"})
	} else {
		checks = append(checks, doctorCheck{Name: "safety.block_destroy_in_prod", Status: doctorWarn, Message: "disabled"})
	}
	if cfg.Policies.RequirePlanFileForApply {
		checks = append(checks, doctorCheck{Name: "safety.require_plan_file_for_apply", Status: doctorPass, Message: "enabled"})
	} else {
		checks = append(checks, doctorCheck{Name: "safety.require_plan_file_for_apply", Status: doctorWarn, Message: "disabled"})
	}
	return checks
}

func checkGit(ctx context.Context, runner doctorCommandRunner) []doctorCheck {
	output, err := runner.Output(ctx, "git", "rev-parse", "--is-inside-work-tree")
	if err != nil || strings.TrimSpace(string(output)) != "true" {
		return []doctorCheck{{Name: "git.repository", Status: doctorWarn, Message: "not inside a Git repository"}}
	}
	checks := []doctorCheck{{Name: "git.repository", Status: doctorPass, Message: "inside a Git repository"}}
	trackedOutput, err := runner.Output(ctx, "git", "ls-files")
	if err != nil {
		checks = append(checks, doctorCheck{Name: "git.tracked_files", Status: doctorWarn, Message: fmt.Sprintf("could not inspect tracked files: %v", err)})
		return checks
	}
	tracked := strings.Fields(string(trackedOutput))
	checks = append(checks, checkTrackedSensitiveFiles(tracked)...)
	return checks
}

func checkTrackedSensitiveFiles(files []string) []doctorCheck {
	tfstate := []string{}
	envFiles := []string{}
	kubeconfigs := []string{}
	for _, file := range files {
		base := strings.ToLower(file)
		switch {
		case strings.HasSuffix(base, ".tfstate") || strings.Contains(base, ".tfstate."):
			tfstate = append(tfstate, file)
		case strings.HasSuffix(base, ".env") || strings.Contains(base, "/.env"):
			envFiles = append(envFiles, file)
		case strings.Contains(base, "kubeconfig") || strings.HasSuffix(base, "/config"):
			kubeconfigs = append(kubeconfigs, file)
		}
	}
	return []doctorCheck{
		trackedFileCheck("git.tracked_tfstate", "tfstate files", tfstate),
		trackedFileCheck("git.tracked_env", ".env files", envFiles),
		trackedFileCheck("git.tracked_kubeconfig", "kubeconfig files", kubeconfigs),
	}
}

func trackedFileCheck(name, label string, files []string) doctorCheck {
	if len(files) == 0 {
		return doctorCheck{Name: name, Status: doctorPass, Message: fmt.Sprintf("no tracked %s", label)}
	}
	sort.Strings(files)
	return doctorCheck{Name: name, Status: doctorWarn, Message: fmt.Sprintf("tracked %s: %s", label, strings.Join(files, ", "))}
}

func printDoctorReport(out interface {
	Write([]byte) (int, error)
}, report doctorReport) {
	fmt.Fprintf(out, "%-6s  %-36s  %s\n", "STATUS", "CHECK", "MESSAGE")
	fmt.Fprintf(out, "%-6s  %-36s  %s\n", "------", "-----", "-------")
	for _, check := range report.Checks {
		fmt.Fprintf(out, "%-6s  %-36s  %s\n", strings.ToUpper(string(check.Status)), check.Name, check.Message)
	}
}

func doctorHasFailure(report doctorReport) bool {
	return report.Status == doctorFail
}

func reportStatus(checks []doctorCheck) doctorStatus {
	status := doctorPass
	for _, check := range checks {
		if check.Status == doctorFail {
			return doctorFail
		}
		if check.Status == doctorWarn {
			status = doctorWarn
		}
	}
	return status
}

func sortedEnvironmentNames(cfg *config.Config) []string {
	names := make([]string, 0, len(cfg.Environments))
	for name := range cfg.Environments {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

var terraformVersionPattern = regexp.MustCompile(`(?m)(?:Terraform|OpenTofu)\s+v?([0-9]+\.[0-9]+\.[0-9]+)`)
var kubernetesVersionPattern = regexp.MustCompile(`(?m)(?:GitVersion|Client Version|Kustomize Version|v)(?:[: ]+)?v?([0-9]+\.[0-9]+\.[0-9]+)`)

func parseTerraformVersion(output string) (string, bool) {
	matches := terraformVersionPattern.FindStringSubmatch(output)
	if len(matches) != 2 {
		return "", false
	}
	return matches[1], true
}

func parseKubernetesVersion(output string) (string, bool) {
	matches := kubernetesVersionPattern.FindStringSubmatch(output)
	if len(matches) != 2 {
		return "", false
	}
	return matches[1], true
}

func compareSemver(left, right string) int {
	l := semverParts(left)
	r := semverParts(right)
	for i := 0; i < 3; i++ {
		if l[i] < r[i] {
			return -1
		}
		if l[i] > r[i] {
			return 1
		}
	}
	return 0
}

func semverParts(version string) [3]int {
	var parts [3]int
	split := strings.Split(version, ".")
	for i := 0; i < len(split) && i < 3; i++ {
		value, _ := strconv.Atoi(split[i])
		parts[i] = value
	}
	return parts
}
