package strings

import "unicode"

// Capitalize capitalizes the first letters of a string.
func Capitalize(s string) string {
	var b []rune
	var upper bool
	for i, r := range s {
		if i == 0 {
			b = append(b, unicode.ToUpper(r))
			continue
		}

		if !unicode.IsLetter(r) {
			upper = true
			b = append(b, r)
			continue
		}

		if upper {
			b = append(b, unicode.ToUpper(r))
			upper = false
		} else {
			b = append(b, unicode.ToLower(r))
		}
	}

	return string(b)
}
