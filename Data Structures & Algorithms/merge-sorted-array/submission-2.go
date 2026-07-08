func merge(nums1 []int, m int, nums2 []int, n int) {
	w := m+n-1
	i, j := m-1, n-1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[w] = nums1[i]
			i--
		} else {
			nums1[w] = nums2[j]
			j--
		}
		w--
	}
}
