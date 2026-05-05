class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        nums_with_index = {}

        for i in range(len(nums)):
            test = target-nums[i]
            if test in nums_with_index and nums_with_index[test] != i:
                return [nums_with_index[test], i]

            nums_with_index[nums[i]] = i