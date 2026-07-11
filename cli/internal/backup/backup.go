package backup

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

const EvidencePath = "backup-evidence.yaml"

type EvidenceFile struct {
	BackupTests map[string]Evidence `yaml:"backup_tests" json:"backup_tests"`
}

type Evidence struct {
	LastBackupTest  string `yaml:"last_backup_test" json:"last_backup_test,omitempty"`
	LastRestoreTest string `yaml:"last_restore_test" json:"last_restore_test,omitempty"`
	Result          string `yaml:"result" json:"result,omitempty"`
	Notes           string `yaml:"notes" json:"notes,omitempty"`
}

type Check struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Report struct {
	Environment string    `json:"environment"`
	GeneratedAt time.Time `json:"generated_at"`
	Checks      []Check   `json:"checks"`
	Evidence    *Evidence `json:"evidence,omitempty"`
	Overall     string    `json:"overall"`
}

func LoadEvidence(path string) (EvidenceFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return EvidenceFile{}, fmt.Errorf("read backup evidence: %w", err)
	}
	var evidence EvidenceFile
	if err := yaml.Unmarshal(data, &evidence); err != nil {
		return EvidenceFile{}, fmt.Errorf("parse backup evidence: %w", err)
	}
	if evidence.BackupTests == nil {
		return EvidenceFile{}, fmt.Errorf("backup evidence must contain backup_tests")
	}
	return evidence, nil
}

func BuildReport(environment, environmentPath, cloud, evidencePath string) (Report, error) {
	text, err := terraformText(environmentPath)
	if err != nil {
		return Report{}, err
	}
	report := Report{Environment: environment, GeneratedAt: time.Now().UTC(), Overall: "passed"}
	report.Checks = append(report.Checks, containsCheck("velero-module", text, "velero", "Velero module/configuration detected.", "Velero module/configuration was not detected."))
	if strings.EqualFold(cloud, "aws") {
		report.Checks = append(report.Checks, containsCheck("backup-bucket", text, "velero-backup-bucket", "AWS backup bucket module detected.", "AWS backup bucket module was not detected."))
	}
	veleroNamespace := strings.Contains(text, `namespace = "velero"`) || strings.Contains(text, `namespace    = "velero"`) || strings.Contains(text, "modules/platform/kubernetes/velero")
	report.Checks = append(report.Checks, booleanCheck("velero-namespace", veleroNamespace, "Velero namespace is expected.", "Expected Velero namespace was not detected."))
	_, runbookErr := os.Stat("docs/backup-restore.md")
	report.Checks = append(report.Checks, booleanCheck("restore-runbook", runbookErr == nil, "Backup and restore runbook exists.", "Backup and restore runbook is missing."))
	evidenceFile, evidenceErr := LoadEvidence(evidencePath)
	if evidenceErr != nil {
		if !os.IsNotExist(unwrapPathError(evidenceErr)) {
			return Report{}, evidenceErr
		}
		report.Checks = append(report.Checks, Check{ID: "restore-evidence", Status: "warn", Message: "Restore test evidence file is missing."})
	} else if item, ok := evidenceFile.BackupTests[environment]; !ok || strings.TrimSpace(item.LastRestoreTest) == "" {
		report.Checks = append(report.Checks, Check{ID: "restore-evidence", Status: "warn", Message: "Restore test evidence is missing for this environment."})
	} else {
		itemCopy := item
		report.Evidence = &itemCopy
		status := "pass"
		message := "Restore test evidence exists."
		if item.Result != "passed" {
			status = "warn"
			message = "Latest recorded backup test did not pass."
		}
		report.Checks = append(report.Checks, Check{ID: "restore-evidence", Status: status, Message: message})
	}
	for _, check := range report.Checks {
		if check.Status == "warn" {
			report.Overall = "warning"
		}
	}
	return report, nil
}

func (report Report) JSON() ([]byte, error) { return json.MarshalIndent(report, "", "  ") }

func Plan(environment string) []string {
	return []string{
		"Review backup scope, retention, encryption, ownership, and the " + environment + " restore runbook.",
		"Create an isolated, non-production restore target with restricted access.",
		"Select a known backup and restore it manually without changing the production source.",
		"Validate objects, persistent data, application health, permissions, and recovery timing.",
		"Destroy the isolated test environment through its reviewed procedure; do not delete the backup.",
		"Record dates, result, non-sensitive notes, approver, and evidence references in backup-evidence.yaml.",
	}
}

func terraformText(root string) (string, error) {
	var result strings.Builder
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			if os.IsNotExist(walkErr) {
				return nil
			}
			return walkErr
		}
		if entry.IsDir() {
			if entry.Name() == ".terraform" {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".tf" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		result.Write(data)
		result.WriteByte('\n')
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("inspect backup Terraform: %w", err)
	}
	return result.String(), nil
}

func containsCheck(id, text, needle, pass, warn string) Check {
	return booleanCheck(id, strings.Contains(strings.ToLower(text), strings.ToLower(needle)), pass, warn)
}
func booleanCheck(id string, ok bool, pass, warn string) Check {
	if ok {
		return Check{ID: id, Status: "pass", Message: pass}
	}
	return Check{ID: id, Status: "warn", Message: warn}
}
func unwrapPathError(err error) error {
	for {
		wrapped, ok := err.(interface{ Unwrap() error })
		if !ok {
			return err
		}
		err = wrapped.Unwrap()
	}
}
