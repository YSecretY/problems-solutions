class Solution:
    def minDeletionSize(self, strs: List[str]) -> int:
        count = 0
        word_size = len(strs[0])
        i = 0
        while i < word_size:
            arr = []
            for j in range(len(strs)):
                arr.append(strs[j][i])
            for k in range(0, len(arr) - 1):
                if arr[k] > arr[k + 1]:
                    count += 1
                    break
            i += 1
        return count
