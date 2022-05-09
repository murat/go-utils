package strings

import "unicode"

// ToKebabCase converts a string to kebab-case.
func ToKebabCase(s string) string {
	b := []rune{}
	for _, r := range s {
		if !unicode.IsLetter(r) && (r == '_' || r == '-' || r == ' ') {
			b = append(b, '-')
			continue
		}

		b = append(b, r)
	}

	return string(b)
}
