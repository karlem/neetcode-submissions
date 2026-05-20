class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        dedup_nums = set(nums)
        best = 0

        for num in dedup_nums:
            if (num - 1) in dedup_nums:
                continue
            
            length = 1
            while (num + length) in dedup_nums:
                length += 1
            
            best = max(best, length)

        return best