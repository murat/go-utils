package strings

import (
	"unicode"
)

// ToCamelCase converts a string to camelCase.
func ToCamelCase(s string) string {
	var b []rune
	var upper bool
	for i, r := range s {
		if i == 0 {
			b = append(b, unicode.ToLower(r))
			continue
		}

		if !unicode.IsLetter(r) {
			upper = true
			continue
		}

		if upper {
			b = append(b, unicode.ToUpper(r))
			upper = false
		} else {
			b = append(b, r)
		}
	}

	return string(b)
}
