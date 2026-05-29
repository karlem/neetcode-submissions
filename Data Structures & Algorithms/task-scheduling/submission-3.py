from collections import defaultdict, deque
from heapq import heapify_max, heappush_max, heappop_max

class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        task_frequency = defaultdict(int)

        for task in tasks:
            task_frequency[task] += 1
        
        h = []
        heapify_max(h)
        q = deque()
        res = []
        time = 0
        
        for task in task_frequency:
            heappush_max(h, (task_frequency[task], task))
        
        while h or q:
            if h:
                task = heappop_max(h)
                if task[0] != 1:
                    q.append((time+n, (task[0] - 1, task[1])))
            if q:
                if q[0][0] == time:
                    task = q.popleft()
                    heappush_max(h, task[1])
        
            time += 1

        return time


        

        