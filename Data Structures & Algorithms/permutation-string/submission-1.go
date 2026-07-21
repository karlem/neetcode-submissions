func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	freq1 := [26]int{}

	for j := 0; j < len(s1); j++ {
		freq1[s1[j]-'a']++
	}

	freq2 := [26]int{}
	for r := 0; r < len(s2); r++ {
		freq2[s2[r]-'a']++

		if r >= len(s1) {
        	freq2[s2[r-len(s1)]-'a']--
    	}

		if freq1 == freq2 {
			return true
		}
	}

	return false
}
