package strings

import (
	"fmt"
	stdStrings "strings"
)

func Pad(s string, length int) string {
	return fmt.Sprintf("%*s", -length, fmt.Sprintf("%*s", (length+len(s))/2, s))
}

func PadWith(s string, length int, pad string) string {
	l := (length - len(s)) / 2
	lpad := stdStrings.Repeat(pad, l)
	rpad := stdStrings.Repeat(pad, length-(l+len(s)))
	return fmt.Sprintf("%*s%s%*s", l, lpad, s, l, rpad)
}

func PadLeft(s string, length int) string {
	return fmt.Sprintf("%*s", length, s)
}

func PadLeftWith(s string, length int, pad string) string {
	l := length - len(s)
	return fmt.Sprintf("%*s%s", l, stdStrings.Repeat(pad, l), s)
}

func PadRight(s string, length int) string {
	return fmt.Sprintf("%s%*s", s, length-len(s), " ")
}

func PadRightWith(s string, length int, pad string) string {
	l := length - len(s)
	return fmt.Sprintf("%s%*s", s, l, stdStrings.Repeat(pad, l))
}
