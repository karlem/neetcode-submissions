func isValid(s string) bool {
    stack := list.New()
	validParentheses := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack.PushBack(c)
			continue
		}

		if stack.Len() == 0 {
			return false
		}
		expected := validParentheses[stack.Remove(stack.Back()).(rune)]
		if expected != c {
			return false
		}
	}

	return stack.Len() == 0
}
