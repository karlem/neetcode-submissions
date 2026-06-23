func isOneEditDistance(s string, t string) bool {
	if len(s) > len(t) {
		s, t = t, s
	}

	if len(t) - len(s) > 1 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			if len(s) == len(t) {
				// If same length, the rest must match perfectly
				return s[i+1:] == t[i+1:]
			} else {
				// If t is longer, s from here must match t from next
				return s[i:] == t[i+1:]
			}
		}
	}

	return len(t) == len(s) + 1
}
