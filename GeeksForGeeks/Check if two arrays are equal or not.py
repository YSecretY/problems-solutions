class Solution:
    def check(self,A,B,N):
        mp = {}
        
        for i in range (N):
            if A[i] in mp.keys ():
                mp[A[i]] += 1
            else:
                mp[A[i]] = 1
        
        for i in range (N):
            if B[i] not in mp.keys ():
                return False
            mp[B[i]] -= 1
        
        for i in mp.keys ():
            if mp[i] != 0:
                return False
                
        return True
