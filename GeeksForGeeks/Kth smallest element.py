class Solution:
    def kthSmallest(self,arr, l, r, k):
        
        hashset = set()
        for n in arr:
            hashset.add(n)
            
        hashset_sorted = sorted(hashset)
        return hashset_sorted[k - 1]
