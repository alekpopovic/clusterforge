package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check local ClusterForge prerequisites",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return err
		}

		ok := true
		for name, engine := range cfg.Engines {
			if _, err := exec.LookPath(engine.Binary); err != nil {
				printer.Warn(fmt.Sprintf("%s binary %q not found", name, engine.Binary))
				ok = false
				continue
			}
			printer.Success(fmt.Sprintf("%s binary found: %s", name, engine.Binary))
		}
		if !ok {
			return fmt.Errorf("one or more checks failed")
		}
		return nil
	},
}
