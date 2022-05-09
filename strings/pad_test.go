package strings_test

import (
	"testing"

	"github.com/murat/go-utils/strings"
)

func TestPad(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path",
			args: args{"hello", 20},
			want: "       hello        ",
		},
		{
			name: "happy path",
			args: args{"hello", 21},
			want: "        hello        ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.Pad(tt.args.s, tt.args.length); got != tt.want {
				t.Errorf("Pad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadWith(t *testing.T) {
	type args struct {
		s      string
		length int
		pad    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path",
			args: args{"hello", 20, "-"},
			want: "-------hello--------",
		},
		{
			name: "happy path",
			args: args{"hello", 21, "-"},
			want: "--------hello--------",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.PadWith(tt.args.s, tt.args.length, tt.args.pad); got != tt.want {
				t.Errorf("PadWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadLeft(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path",
			args: args{"hello", 20},
			want: "               hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.PadLeft(tt.args.s, tt.args.length); got != tt.want {
				t.Errorf("PadLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadLeftWith(t *testing.T) {
	type args struct {
		s      string
		length int
		pad    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path",
			args: args{"hello", 20, "-"},
			want: "---------------hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.PadLeftWith(tt.args.s, tt.args.length, tt.args.pad); got != tt.want {
				t.Errorf("PadLeftWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadRight(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path",
			args: args{"hello", 20},
			want: "hello               ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.PadRight(tt.args.s, tt.args.length); got != tt.want {
				t.Errorf("PadRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadRightWith(t *testing.T) {
	type args struct {
		s      string
		length int
		pad    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path",
			args: args{"hello", 20, "-"},
			want: "hello---------------",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.PadRightWith(tt.args.s, tt.args.length, tt.args.pad); got != tt.want {
				t.Errorf("PadRightWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
