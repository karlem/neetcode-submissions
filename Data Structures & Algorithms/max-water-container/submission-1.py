class Solution:
    def maxArea(self, heights: List[int]) -> int:
        # best = max(best, (right - left) * min(heights[left], heights[right]))
        left, right = 0, len(heights) - 1
        best = 0

        while left < right:
            best = max(best, (right - left) * min(heights[left], heights[right]))

            if heights[left] < heights[right]:
                left += 1
            else:
                right -= 1

        return best