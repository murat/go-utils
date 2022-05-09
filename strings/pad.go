package strings

import (
	"fmt"
	stdStrings "strings"
)

// Pad centers a string into given length.
func Pad(s string, length int) string {
	return fmt.Sprintf("%*s", -length, fmt.Sprintf("%*s", (length+len(s))/2, s))
}

// PadWith centers a string into given length with given pad.
func PadWith(s string, length int, pad string) string {
	l := (length - len(s)) / 2
	lpad := stdStrings.Repeat(pad, l)
	rpad := stdStrings.Repeat(pad, length-(l+len(s)))
	return fmt.Sprintf("%*s%s%*s", l, lpad, s, l, rpad)
}

// PadLeft adds padding to the left of a string.
func PadLeft(s string, length int) string {
	return fmt.Sprintf("%*s", length, s)
}

// PadLeftWith adds padding to the left of a string with given pad.
func PadLeftWith(s string, length int, pad string) string {
	l := length - len(s)
	return fmt.Sprintf("%*s%s", l, stdStrings.Repeat(pad, l), s)
}

// PadRight adds padding to the right of a string.
func PadRight(s string, length int) string {
	return fmt.Sprintf("%s%*s", s, length-len(s), " ")
}

// PadRightWith adds padding to the right of a string with given pad.
func PadRightWith(s string, length int, pad string) string {
	l := length - len(s)
	return fmt.Sprintf("%s%*s", s, l, stdStrings.Repeat(pad, l))
}
