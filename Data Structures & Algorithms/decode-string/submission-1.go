// 2[a3[b]]c

func decodeString(s string) string {
	i := 0

	var decode func() string
	decode = func() string {
		res := ""
		k := 0

		for i < len(s) {
			c := s[i]

			if c >= '0' && c <= '9' {
				k = k*10+int(c-'0')
			} else if c == '[' {
				i++
				innerRes := decode()
				for j := 0; j < k; j++ {
					res += innerRes
				}
				k = 0
			} else if c == ']' {
				return res
			} else {
				res += string(c)
			}

			i++
		}

		return res
	}

	return decode()
}
