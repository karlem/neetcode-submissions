class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        rows = len(matrix)
        cols = len(matrix[0])

        if rows == 0 or cols == 0:
            return False

        left, right = 0, rows * cols - 1

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