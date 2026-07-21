func decodeString(s string) string {
	i := 0
	var rec func() string
	rec = func() string {
		var count int
		res := ""

		for i < len(s) {
			if s[i] >= '0' && s[i] <= '9' {
				n := int(s[i]-'0')
				count *=10
				count += n
			} else if s[i] == '[' {
				i++
				
				inRes := rec()
				for j := 0; j < count; j++ {
					res += inRes
				}
				count = 0
			} else if s[i] == ']' {
				break
			} else {
				res += string(s[i])
			}
			
			i++
		}
		return res
	}

	return rec()
}
