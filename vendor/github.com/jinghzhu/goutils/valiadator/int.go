package valiadator

// IsInt check if the string is an integer. Empty string is valid.
func IsInt(str string) bool {
	if IsNull(str) {
		return true
	}

	return rxInt.MatchString(str)
}
