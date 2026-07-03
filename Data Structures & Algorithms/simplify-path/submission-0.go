func simplifyPath(path string) string {
	currentPath := []string{}

	i := 0
	for i < len(path) {
		if path[i] == '/' {
			op := ""
			j := i+1
			for j < len(path) && path[j] != '/' {
				op += string(path[j])
				j++
			}
			
			if op == "." || op == "" {
				// no op
			} else if op == ".." {
				if len(currentPath) > 0 {
					currentPath = currentPath[:len(currentPath)-1]
				}
			} else {
				currentPath = append(currentPath, op)
			}

			i = j
		}
	}

	res := "/"
	for i, s := range currentPath {
		if i != 0 {
			res += "/"
		}
		res += s
	}

	return res
}
