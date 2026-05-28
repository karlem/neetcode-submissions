import heapq

class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        max_heap = []

        for num in nums:
            heapq.heappush(max_heap, -num)
        
        kth_largest = 0
        for i in range(k):
            kth_largest = -heapq.heappop(max_heap)

        return kth_largest