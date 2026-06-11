package planjson

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

type Summary struct {
	Creates      int
	Updates      int
	Deletes      int
	Replacements int
	NoOps        int
	Addresses    []string
}

type plan struct {
	ResourceChanges []resourceChange `json:"resource_changes"`
}

type resourceChange struct {
	Address string `json:"address"`
	Change  change `json:"change"`
}

type change struct {
	Actions []string `json:"actions"`
}

func Parse(data []byte) (Summary, error) {
	var plan plan
	if err := json.Unmarshal(data, &plan); err != nil {
		return Summary{}, fmt.Errorf("parse Terraform plan JSON: %w", err)
	}

	var summary Summary
	for _, resource := range plan.ResourceChanges {
		actions := actionSet(resource.Change.Actions)
		switch {
		case actions["no-op"]:
			summary.NoOps++
		case actions["delete"] && actions["create"]:
			summary.Replacements++
			summary.Addresses = append(summary.Addresses, resource.Address)
		case actions["delete"]:
			summary.Deletes++
			summary.Addresses = append(summary.Addresses, resource.Address)
		case actions["create"]:
			summary.Creates++
			summary.Addresses = append(summary.Addresses, resource.Address)
		case actions["update"]:
			summary.Updates++
			summary.Addresses = append(summary.Addresses, resource.Address)
		}
	}
	sort.Strings(summary.Addresses)
	if len(summary.Addresses) > 20 {
		summary.Addresses = summary.Addresses[:20]
	}
	return summary, nil
}

func Print(w io.Writer, env string, summary Summary, risk string, policy string) {
	fmt.Fprintf(w, "Plan summary for %s:\n", env)
	fmt.Fprintf(w, "  Create:  %d\n", summary.Creates)
	fmt.Fprintf(w, "  Update:  %d\n", summary.Updates)
	fmt.Fprintf(w, "  Delete:  %d\n", summary.Deletes)
	fmt.Fprintf(w, "  Replace: %d\n", summary.Replacements)
	fmt.Fprintf(w, "  No-op:   %d\n", summary.NoOps)
	fmt.Fprintf(w, "Risk: %s\n", risk)
	if policy != "" {
		fmt.Fprintf(w, "Policy: %s\n", policy)
	}
	if len(summary.Addresses) > 0 {
		fmt.Fprintln(w, "Changed resources:")
		for _, address := range summary.Addresses {
			fmt.Fprintf(w, "  - %s\n", address)
		}
	}
}

func actionSet(actions []string) map[string]bool {
	set := make(map[string]bool, len(actions))
	for _, action := range actions {
		set[action] = true
	}
	return set
}
