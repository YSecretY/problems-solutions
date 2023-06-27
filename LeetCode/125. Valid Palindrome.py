class Solution:
    def isPalindrome(self, s: str) -> bool:
        res = ""
        s = s.lower()
        for i in range(len(s)):
            if s[i].isalpha() or s[i].isdigit():
                res += s[i]

        print(res)

        for i in range(len(res)):
            if res[-i-1] != res[i]:
                return False

        return True
