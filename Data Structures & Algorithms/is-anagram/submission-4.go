func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := make([]rune, 26)
	for i := range s {
		freq[t[i]-'a']++
		freq[s[i]-'a']--
	}

    for _, val := range freq {
        if val != 0 {
            return false
        }
    }

    return true
}
