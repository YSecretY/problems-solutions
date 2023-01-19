class Solution:
    def reverseWords(self, s):
        words = s.split('.')
        words.reverse()
        res = '.'.join(words)
        return res
