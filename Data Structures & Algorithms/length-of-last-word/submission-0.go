func lengthOfLastWord(s string) int {
	split := strings.Split(strings.TrimSpace(s), " ")

	return len(split[len(split)-1])
}
