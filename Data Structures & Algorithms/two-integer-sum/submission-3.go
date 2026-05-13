func twoSum(nums []int, target int) []int {
    seen := map[int]int{}

    for i, num := range nums {
        if foundIndex, ok := seen[target - num]; ok {
            return []int{foundIndex, i}
        }

        seen[num] = i
    }

    return []int{}
} 