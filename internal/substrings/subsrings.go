package substrings

func IsSubstring(mainstr, substr string) bool {
	for i := 0; i <= len(mainstr)-len(substr); i++ {
		found := true
		for j := 0; j < len(substr); j++ {
			if mainstr[i+j] != substr[j] {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}
	return false
}
