class Solution:
    def totalHours(self, piles: List[int], rate: int) -> int:
        total = 0
        for pile in piles:
            total += math.ceil(pile/rate)
        return total

    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        maxPile, minPile = 0, 1

        for pile in piles:
            maxPile = max(maxPile, pile)

        while minPile < maxPile:
            mid = (minPile + maxPile) // 2
            totalHours = self.totalHours(piles, mid)

            if totalHours <= h:
                maxPile = mid
            else:
                minPile = mid + 1

        return minPile
