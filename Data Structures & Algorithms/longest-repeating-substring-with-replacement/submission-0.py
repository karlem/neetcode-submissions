class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        left = 0
        longest = 0
        freq = collections.defaultdict(int)
        most_frequent = 0

        for right in range(len(s)):
            freq[s[right]] += 1
            most_frequent = max(most_frequent, freq[s[right]])

            while ((right - left + 1) - most_frequent) > k:
                freq[s[left]] -= 1
                left += 1


            longest = max(longest, right - left + 1)
        
        return longest