class Solution:
    def countBits(self, n: int) -> List[int]:
        counts = []

        for i in range(n+1):
            count = 0

            for j in range(i):
                if (i >> j) & 1:
                    count += 1;
            
            counts.append(count)

        return counts