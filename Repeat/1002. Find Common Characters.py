class Solution:
    def commonChars(self, words: List[str]) -> List[str]:
        res = list(words[0])
        for word in words:
            check = []
            for letter in word:
                if letter in res:
                    check.append(letter)
                    res.remove(letter)
            res = check

        return res
