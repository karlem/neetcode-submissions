func simplifyPath(path string) string {
	stack := []string{}
	paths := strings.Split(path, "/")

	for _, curr := range paths {
		if curr == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if curr != "." && curr != ""  {
			stack = append(stack, curr)
		}
	}

	return "/" + strings.Join(stack, "/")
}
