func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := range strs {
		for j := 0; j < len(prefix); j++ {
			if j >= len(strs[i]) || prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
	}

	return prefix
}
