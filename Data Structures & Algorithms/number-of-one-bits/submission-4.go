func hammingWeight(n int) int {
	var count int
	for i := 0; i < 31; i++ {
		if n & (1<<i) != 0 {
			count++
		}
	}
	return count
}
