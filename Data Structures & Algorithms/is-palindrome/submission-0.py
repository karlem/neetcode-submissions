class Solution:
    def isPalindrome(self, s: str) -> bool:
        left, right = 0, len(s) - 1

        while left < right:
            if not self.isAlphaNum(s[left]):
                left += 1
                continue

            if not self.isAlphaNum(s[right]):
                right -= 1
                continue
            
            if s[left].lower() != s[right].lower():
                return False
            
            left += 1
            right -= 1
        
        return True
    
    def isAlphaNum(self, char: str):
	    return char.isalpha() or char.isdigit()