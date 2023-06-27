class Solution:
    def convBinToDec(self, val: str) -> int:
        res, i = 0, 0
        for dig in val[::-1]:
            res += int(int(dig) * 2**i)
            i += 1
        return res

    def convDecToBin(self, val: int) -> str:
        res = ""
        while val > 0:
            res += str(int(val % 2))
            val //= 2
        return res[::-1] if res != "" else "0"

    def addBinary(self, a: str, b: str) -> str:
        a_dec = self.convBinToDec(a)
        b_dec = self.convBinToDec(b)
        return self.convDecToBin(a_dec + b_dec)
