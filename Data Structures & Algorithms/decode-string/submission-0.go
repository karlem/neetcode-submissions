func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func decodeString(s string) string {
	i := 0

	// 20[a3[b]]c
	var decode func() string
	decode = func() string {
		res := ""
		k := 0

		for i < len(s) {
			b := s[i]

			switch {
				case isDigit(b):
					k = k*10 + int(b-'0')
				case b == '[':
					i++
					subRes := decode()
					for j := 0; j < k; j++ {
						res += subRes
					}
					k = 0
				case b == ']':
					return res
				default:
					res += string(b)
			}

			i++
		}

		return res
	}

	return decode()
}