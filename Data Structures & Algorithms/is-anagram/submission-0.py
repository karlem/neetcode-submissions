class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        
        char_count = {}

        for c in s:
            if c in char_count:
                char_count[c] += 1
            else:
                char_count[c] = 1
        
        for c in t:
            if c not in char_count:
                return False

            if char_count[c] < 1:
                return False
            
            char_count[c] -= 1
        
        return True

