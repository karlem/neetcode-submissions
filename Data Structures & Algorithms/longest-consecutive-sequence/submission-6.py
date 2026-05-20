class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        dedup_nums = set(nums)
        if not dedup_nums:
            return 0

        best = 1

        for num in dedup_nums:
            if num-1 in dedup_nums:
                continue
            
            i = 1
            while num + i in dedup_nums:
                best = max(best, i+1)
                i += 1
            
        return best