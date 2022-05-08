package strings_test

import (
	"testing"

	"github.com/murat/go-utils/strings"
)

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "hello world",
			output: "Hello World",
		},
		{
			input:  "HEllo WOrld",
			output: "Hello World",
		},
		{
			input:  "hello-world",
			output: "Hello-World",
		},
		{
			input:  "hello__world",
			output: "Hello__World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := strings.Capitalize(tt.input); got != tt.output {
				t.Errorf("Capitalize() = %v, want %v", got, tt.output)
			}
		})
	}
}
