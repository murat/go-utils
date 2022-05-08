package strings_test

import (
	"testing"

	"github.com/murat/go-utils/strings"
)

func TestStartsWith(t *testing.T) {
	type args struct {
		s      string
		starts string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy",
			args: args{
				s:      "hello",
				starts: "hel",
			},
			want: true,
		},
		{
			name: "happy",
			args: args{
				s:      "happy",
				starts: "hap",
			},
			want: true,
		},
		{
			name: "unhappy",
			args: args{
				s:      "unhappy",
				starts: "happy",
			},
			want: false,
		},
		{
			name: "unhappy",
			args: args{
				s:      "unhappy",
				starts: "unhappybutmore",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.StartsWith(tt.args.s, tt.args.starts); got != tt.want {
				t.Errorf("StartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
