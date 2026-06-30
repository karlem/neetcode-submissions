// ["t","h","e"," ","s","k","y"," ","i","s"," ","b","l","u","e"]
// ["e","u","l","b"," ","s","i"," ","y","k","s"," ","e","h","t"]
// ["b","l","u","e"," ","s","i"," ","y","k","s"," ","e","h","t"]

func reverseWords(s []byte) {
	left, right := 0, len(s) - 1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	l := 0
	for l < len(s) {
		r := l
		for r < len(s) && s[r] != ' ' {
			r++
		}

		wl, wr := l, r-1
		for wl < wr {
			s[wl], s[wr] = s[wr], s[wl]
			wl++
			wr--
		}

		l = r+1
	}
}
