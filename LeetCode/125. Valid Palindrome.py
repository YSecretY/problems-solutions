class Solution:
    def isPalindrome(self, s: str) -> bool:
        res = [char for _, char in enumerate(s.lower()) if char.isalpha() or char.isdigit()]
        return res == res[::-1]
