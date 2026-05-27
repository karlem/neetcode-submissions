from collections import defaultdict

class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        left = 0
        freq = defaultdict(int)
        mostFrequentChar = 0
        longestRepeating = 0

        for right in range(len(s)):
            freq[s[right]] += 1
            mostFrequentChar = max(mostFrequentChar, freq[s[right]])
            
            while ((right - left + 1) - mostFrequentChar) > k:
                freq[s[left]] -= 1
                left += 1

            longestRepeating = max(longestRepeating, right - left + 1)

        return longestRepeating