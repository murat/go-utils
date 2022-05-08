package strings

import "testing"

func TestKebabCase(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "hello world",
			output: "hello-world",
		},
		{
			input:  "Hello_World",
			output: "Hello-World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := KebabCase(tt.input); got != tt.output {
				t.Errorf("KebabCase() = %v, want %v", got, tt.output)
			}
		})
	}
}
