package environment

import "strings"

func IsProduction(name string) bool {
	normalized := strings.ToLower(strings.TrimSpace(name))
	if normalized == "" {
		return false
	}

	tokens := strings.FieldsFunc(normalized, func(r rune) bool {
		switch r {
		case '-', '_', '.', '/', '\\', ':':
			return true
		default:
			return false
		}
	})
	if len(tokens) == 0 {
		return false
	}

	switch tokens[0] {
	case "prod", "production", "prd":
		return true
	default:
		return false
	}
}
