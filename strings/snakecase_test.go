package strings_test

import (
	"testing"

	"github.com/murat/go-utils/strings"
)

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "hello world",
			output: "hello_world",
		},
		{
			input:  "Hello world",
			output: "hello_world",
		},
		{
			input:  "Hello-World",
			output: "hello_world",
		},
		{
			input:  "Hello!World",
			output: "hello_world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := strings.ToSnakeCase(tt.input); got != tt.output {
				t.Errorf("ToSnakeCase() = %v, want %v", got, tt.output)
			}
		})
	}
}
