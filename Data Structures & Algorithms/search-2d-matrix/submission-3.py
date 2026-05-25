# 3 x 4
# (rows * cols) - 1

# 0
# 11 = (rows * cols) - 1

# (left + right) // 2 = (0 + 11) // 2 = 5
# i = mid // cols
# j = mid % cols

class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        cols = len(matrix[0])

        if len(matrix) == 0 or cols == 0:
            return False

        left, right = 0, (len(matrix) * cols) - 1

        while left <= right:
            mid = (left + right) // 2
            i = mid // cols
            j = mid % cols

            if matrix[i][j] == target:
                return True
            
            if matrix[i][j] < target:
                left = mid + 1
            else:
                right = mid - 1

        return False