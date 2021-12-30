package utils

func Contains(target []string, contains string) bool {
	for _, entry := range target {
		if entry == contains {
			return true
		}
	}
	return false
}
