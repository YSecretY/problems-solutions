class Solution:    
    def doUnion(self,a,n,b,m):
        hashset = set()
        size = 0
        
        for i in a:
            if i not in hashset:
                hashset.add(i)
                size += 1
        
        for i in b:
            if i not in hashset:
                hashset.add(i)
                size += 1
                
        return size
