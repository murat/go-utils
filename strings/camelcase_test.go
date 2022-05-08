package strings_test

import (
	"testing"

	"github.com/murat/go-utils/strings"
)

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "hello world",
			output: "helloWorld",
		},
		{
			input:  "hello--world",
			output: "helloWorld",
		},
		{
			input:  "HelloWorld",
			output: "helloWorld",
		},
		{
			input:  "hello-_-world",
			output: "helloWorld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := strings.ToCamelCase(tt.input); got != tt.output {
				t.Errorf("ToCamelCase() = %v, want %v", got, tt.output)
			}
		})
	}
}
