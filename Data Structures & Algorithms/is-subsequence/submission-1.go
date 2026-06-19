func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	var j int
	for i := 0; i < len(s); i++ {
		for j <= len(t) {
			if j == len(t) {
				return false
			}

			if s[i] == t[j] {
				j++
				break
			}

			j++
		}
	}

	return true
}
