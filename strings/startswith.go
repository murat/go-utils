package strings

func StartsWith(s, starts string) bool {
	if len(starts) > len(s) {
		return false
	}

	b := []byte(s)
	for i, r := range starts {
		if rune(b[i]) != r {
			return false
		}
	}

	return true
}
