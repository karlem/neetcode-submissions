func calPoints(operations []string) int {
	stack := []int{}

	for _, op := range operations {
		switch op {
			case "+":
				l := len(stack)
				stack = append(stack, stack[l-1] + stack[l-2])
			case "D":
				l := len(stack)
				stack = append(stack, stack[l-1]*2)
			case "C":
				l := len(stack)
				stack = stack[:l-1]
			default:
				num, _ := strconv.Atoi(op)
				stack = append(stack, num)
		}
	}

	var sum int
	for _, num := range stack {
		sum += num
	}
	return sum
}
