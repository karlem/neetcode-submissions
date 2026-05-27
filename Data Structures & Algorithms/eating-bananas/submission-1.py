class Solution:
    def totalHours(self, piles: List[int], rate: int) -> int:
        total = 0
        for pile in piles:
            total += (pile + rate - 1) // rate
        return total

    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        left, right = 1, max(piles)

        while left < right:
            mid = (left + right) // 2
            if self.totalHours(piles, mid) <= h:
                right = mid
            else:
                left = mid + 1

        return left