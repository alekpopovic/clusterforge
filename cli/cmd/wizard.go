package cmd

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type promptSession struct {
	reader *bufio.Reader
	out    io.Writer
}

func newPromptSession(cmd *cobra.Command) *promptSession {
	return &promptSession{
		reader: bufio.NewReader(cmd.InOrStdin()),
		out:    cmd.OutOrStdout(),
	}
}

func optionalArg(args []string, index int) string {
	if len(args) <= index {
		return ""
	}
	return strings.TrimSpace(args[index])
}

func requireValue(cmd *cobra.Command, value, label string) (string, error) {
	return requireValueWithPrompt(value, label, newPromptSession(cmd))
}

func requireValueWithPrompt(value, label string, prompts *promptSession) (string, error) {
	value = strings.TrimSpace(value)
	if value != "" {
		return value, nil
	}
	if opts.NonInteractive {
		return "", fmt.Errorf("%s is required in non-interactive mode", label)
	}
	return prompts.String(label, "")
}

func promptString(in io.Reader, out io.Writer, label, defaultValue string) (string, error) {
	return (&promptSession{reader: bufio.NewReader(in), out: out}).String(label, defaultValue)
}

func (p *promptSession) String(label, defaultValue string) (string, error) {
	if defaultValue == "" {
		fmt.Fprintf(p.out, "%s: ", label)
	} else {
		fmt.Fprintf(p.out, "%s [%s]: ", label, defaultValue)
	}
	value, err := p.reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	value = strings.TrimSpace(value)
	if value == "" {
		return defaultValue, nil
	}
	return value, nil
}

func promptInt(in io.Reader, out io.Writer, label string, defaultValue int) (int, error) {
	return (&promptSession{reader: bufio.NewReader(in), out: out}).Int(label, defaultValue)
}

func (p *promptSession) Int(label string, defaultValue int) (int, error) {
	value, err := p.String(label, strconv.Itoa(defaultValue))
	if err != nil {
		return 0, err
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a number", label)
	}
	return parsed, nil
}

func promptBool(in io.Reader, out io.Writer, label string, defaultValue bool) (bool, error) {
	return (&promptSession{reader: bufio.NewReader(in), out: out}).Bool(label, defaultValue)
}

func (p *promptSession) Bool(label string, defaultValue bool) (bool, error) {
	defaultText := "n"
	if defaultValue {
		defaultText = "y"
	}
	value, err := p.String(label+" (y/n)", defaultText)
	if err != nil {
		return false, err
	}
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "y", "yes", "true":
		return true, nil
	case "n", "no", "false":
		return false, nil
	default:
		return false, fmt.Errorf("%s must be yes or no", label)
	}
}
