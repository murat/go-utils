package slices_test

import (
	"testing"

	"github.com/murat/go-utils/slices"
)

// func TestContains(t *testing.T) {
// 	intSlice := []int{1, 2, 3, 4, 5}
// 	if !slices.Contains(intSlice, 2) {
// 		t.Errorf("expected slice to contain 2")
// 	}
// 	if slices.Contains(intSlice, 6) {
// 		t.Errorf("expected slice not to contain 6")
// 	}

// 	strSlice := []string{"hello", "world"}
// 	if !slices.Contains(strSlice, "hello") {
// 		t.Errorf("expected slice to contain hello")
// 	}
// 	if slices.Contains(strSlice, "murat") {
// 		t.Errorf("expected slice not to contain murat")
// 	}
// }

func TestContains(t *testing.T) {
	type args[T comparable] struct {
		slice []T
		val   T
	}

	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}

	for _, tt := range []testCase[int]{
		{"happy ints path", args[int]{[]int{1, 2, 3}, 1}, true},
		{"unhappy ints path", args[int]{[]int{1, 2, 3}, 4}, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.Contains(tt.args.slice, tt.args.val); got != tt.want {
				t.Errorf("Contains(%v, %v) = %v, want %v", tt.args.slice, tt.args.val, got, tt.want)
			}
		})
	}

	for _, tt := range []testCase[string]{
		{"happy strings path", args[string]{[]string{"1", "2", "3"}, "1"}, true},
		{"unhappy strings path", args[string]{[]string{"1", "2", "3"}, "4"}, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.Contains(tt.args.slice, tt.args.val); got != tt.want {
				t.Errorf("Contains(%v, %v) = %v, want %v", tt.args.slice, tt.args.val, got, tt.want)
			}
		})
	}
}
