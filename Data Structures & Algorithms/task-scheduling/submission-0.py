from collections import defaultdict, deque
from heapq import heapify_max, heappush_max, heappop_max

class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        # this could be Counter from collections
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
        
        # look at the stack and see if there is something to run
        # if there is then run it and after running it put in on queue with
        # current + n.
        # Check if queue has someting ready to be put back on heap
        # (time == time to run) and if yes then put it there.
        # If heap is empty but queue not then just add "Idle" to result.
        while h or q:
            if h:
                task = heappop_max(h)
                res.append(task)

                if task[0] != 1:
                    q.append((time+n, (task[0] - 1, task[1])))
            else:
                res.append("Idle")
            
            if q:
                if q[0][0] == time:
                    task = q.popleft()
                    heappush_max(h, task[1])
        
            time += 1

        print(res)

        return len(res)


        

        