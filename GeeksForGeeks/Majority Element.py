class Solution:
    def majorityElement(self, A, N):
        mp = {}
        
        for number in A:
            mp.setdefault(number, 0)
            mp[number] += 1
            if mp[number] > N / 2:
                return number
                
        return -1
