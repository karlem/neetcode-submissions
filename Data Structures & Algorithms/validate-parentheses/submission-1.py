class Solution:
    def isValid(self, s: str) -> bool:
        openBrackets = []

        for c in s:
            if self.isOpenBracket(c):
                openBrackets.append(c)
                continue
            
            if len(openBrackets) == 0:
                return False

            lastOpen = openBrackets.pop()
            if lastOpen == '(' and c != ')':
                return False
            
            if lastOpen == '[' and c != ']':
                return False
            
            if lastOpen == '{' and c != '}':
                return False
        
        return len(openBrackets) == 0

            
            
    def isOpenBracket(self, c: str) -> bool:
        return c == '(' or c == '[' or c == '{'