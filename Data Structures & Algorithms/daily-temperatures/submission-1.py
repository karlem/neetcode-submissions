class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        stack = []
        output = [0] * len(temperatures)

        for currentIndex, currentTemp in enumerate(temperatures):
            while len(stack) and stack[-1][1] < currentTemp:
                prevIndex, prevTemp = stack.pop()
                output[prevIndex] = currentIndex - prevIndex

            stack.append((currentIndex, currentTemp))
        
        return output