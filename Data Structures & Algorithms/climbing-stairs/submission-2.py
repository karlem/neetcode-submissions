class Solution:
    def climbStairs(self, n: int) -> int:
        mem = [-1] * (n+1)

        def dfs(n: int):
            if mem[n] != -1:
                return mem[n]

            if n <= 2:
                return n

            mem[n] = dfs(n - 1) + dfs(n - 2)
            return mem[n]
        
        return dfs(n)
