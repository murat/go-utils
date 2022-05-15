package slices

import "testing"

func TestContains(t *testing.T) {
	type args struct {
		slice []interface{}
		val   interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns true",
			args: args{
				slice: []interface{}{"a", "b", 1},
				val:   1,
			},
			want: true,
		},
		{
			name: "returns false",
			args: args{
				slice: []interface{}{"a", "b", 1},
				val:   2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.slice, tt.args.val); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
