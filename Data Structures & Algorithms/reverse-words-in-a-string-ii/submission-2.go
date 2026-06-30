func reverse(s []byte, left int, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func reverseWords(s []byte) {
	reverse(s, 0, len(s)-1)

	l := 0
	for l < len(s) {
		r := l
		for r < len(s) && s[r] != ' ' {
			r++
		}

		reverse(s, l, r-1)
		l = r+1
	}
}
