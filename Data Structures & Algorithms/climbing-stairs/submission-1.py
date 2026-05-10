class Solution:
    def climbStairs(self, n: int) -> int:
        mem = {}
        return self.climbStairsRec(n, mem)

    def climbStairsRec(self, n: int, mem):
        if n in mem:
            return mem[n]
            
        if n < 1:
            return 0
        if n == 1:
            return 1
        if n == 2:
            return 2

        mem[n] = self.climbStairsRec(n - 1, mem) + self.climbStairsRec(n - 2, mem)
        return mem[n]
