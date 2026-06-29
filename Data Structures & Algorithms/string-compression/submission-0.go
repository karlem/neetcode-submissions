func compress(chars []byte) int {
    l, w := 0, 0

	for l < len(chars) {
		r := l
		for r < len(chars) && chars[l] == chars[r] {
			r++
		}

		chars[w] = chars[l]
		w++

		count := r-l
		if count > 1 {
			countStr := strconv.Itoa(count)
			for i := 0; i < len(countStr); i++ {
				chars[w] = countStr[i]
				w++
			}
		}

		l = r
	}

	return w
}
