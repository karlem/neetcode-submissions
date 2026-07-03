func compress(chars []byte) int {
    i, w := 0, 0
	for i < len(chars) {
		j := i
		for j < len(chars) && chars[i] == chars[j] {
			j++
		}

		chars[w] = chars[i]
		w++

		count := j-i
		if count > 1 {
			s := strconv.Itoa(count)
			for i := 0; i < len(s); i++ {
				chars[w] = s[i]
				w++
			}
		}

		i = j
	}

	return w
}
