from collections import defaultdict, deque
from heapq import heapify_max, heappush_max, heappop_max

class Solution:
     def leastInterval(self, tasks, n):
        count = Counter(tasks)
        heap = [-c for c in count.values()]
        heapq.heapify(heap)

        cooldown = deque()  # (ready_time, count)
        time = 0

        while heap or cooldown:
            time += 1

            # release tasks whose cooldown is done
            while cooldown and cooldown[0][0] == time:
                _, cnt = cooldown.popleft()
                heapq.heappush(heap, cnt)

            if heap:
                cnt = heapq.heappop(heap) + 1  # negative counts
                if cnt != 0:
                    cooldown.append((time + n + 1, cnt))

        return time

        

        