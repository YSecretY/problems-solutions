class Solution:
    def leaders(self, A, N):
        max = 0
        res = []
        for i in range(N - 1, -1, -1):
            if A[i] >= max:
                res.append(A[i])
                max = A[i]
        return res[::-1]
