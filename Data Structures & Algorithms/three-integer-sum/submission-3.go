import "slices"
//  [-4,-1,-1,0,1,2]
func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	res := [][]int{}
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums) - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i],nums[j],nums[k]})
				
				j++
				k--

				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return res
}
