package ui

import (
	"encoding/json"
	"fmt"
	"io"
)

type Printer struct {
	out io.Writer
	err io.Writer
}

func NewPrinter(out, err io.Writer) Printer {
	return Printer{out: out, err: err}
}

func (p Printer) Info(message string) {
	fmt.Fprintln(p.out, message)
}

func (p Printer) Success(message string) {
	fmt.Fprintf(p.out, "ok: %s\n", message)
}

func (p Printer) Warn(message string) {
	fmt.Fprintf(p.err, "warning: %s\n", message)
}

func WriteJSON(out io.Writer, value any) error {
	encoder := json.NewEncoder(out)
	encoder.SetIndent("", "  ")
	return encoder.Encode(value)
}
