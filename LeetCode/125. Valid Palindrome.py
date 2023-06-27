class Solution:
    def isPalindrome(self, s: str) -> bool:
        res = ""
        s = s.lower()
        for i in range(len(s)):
            if s[i].isalpha() or s[i].isdigit():
                res += s[i]

        return res == res[::-1]
