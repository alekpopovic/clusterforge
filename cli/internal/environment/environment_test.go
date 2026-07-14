package environment

import "testing"

func TestIsProduction(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{name: "prod", want: true},
		{name: "production", want: true},
		{name: "prd", want: true},
		{name: "prod-eu", want: true},
		{name: "prod_eu", want: true},
		{name: "production.eu", want: true},
		{name: "prd/eu", want: true},
		{name: " Prod-US ", want: true},
		{name: "dev", want: false},
		{name: "staging", want: false},
		{name: "product", want: false},
		{name: "sandbox-prod", want: false},
		{name: "---", want: false},
		{name: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsProduction(tt.name); got != tt.want {
				t.Fatalf("IsProduction(%q) = %t, want %t", tt.name, got, tt.want)
			}
		})
	}
}
