package slices

// Contains returns true if the given slice contains the given value.
func Contains(slice []interface{}, val interface{}) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
