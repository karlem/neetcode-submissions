class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        dedup_nums = set(nums)
        if not dedup_nums:
            return 0

        best = 1

        for num in dedup_nums:
            if (num-1) in dedup_nums:
                continue
            
            lenght = 1
            while num + lenght in dedup_nums:
                best = max(best, lenght+1)
                lenght += 1
            
        return best