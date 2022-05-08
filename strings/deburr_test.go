package strings_test

import (
	"testing"

	"github.com/murat/go-utils/strings"
)

func TestDeburr(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "héllo",
			output: "hello",
		},
		{
			input:  "déjà vu",
			output: "deja vu",
		},
		{
			input:  "résumé",
			output: "resume",
		},
		{
			input:  "žůžo",
			output: "zuzo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := strings.Deburr(tt.input); got != tt.output {
				t.Errorf("Deburr() = %v, want %v", got, tt.output)
			}
		})
	}
}
