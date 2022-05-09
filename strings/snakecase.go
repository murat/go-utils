package strings

import (
	"strings"
	"unicode"
)

// ToSnakeCase converts a string to snake_case.
func ToSnakeCase(s string) string {
	var b []rune
	for _, r := range strings.ToLower(s) {
		if !unicode.IsLetter(r) || r == '_' {
			b = append(b, '_')
			continue
		}
		b = append(b, unicode.ToLower(r))
	}

	return string(b)
}
