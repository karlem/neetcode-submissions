func wordBreak(s string, wordDict []string) bool {
	words := map[string]struct{}{}
	for _, word := range wordDict {
		words[word] = struct{}{}
	}

	mem := map[int]bool{}

	var canBreak func (int) bool
	canBreak = func (i int) bool {
		if i == len(s) {
			return true
		}

		if _, ok := mem[i]; ok {
			return mem[i]
		}
		
		for j := i+1; j <= len(s); j++ {
			_, wordMatched := words[s[i:j]]
			if wordMatched && canBreak(j) {
				mem[i] = true
				return true
			}
		}

		mem[i] = false
		return false
	}

	
	return canBreak(0)
}
