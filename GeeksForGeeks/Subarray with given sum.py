class Solution:
    def subArraySum(self,arr, n, s):
        cur_sum = arr[0]
        left = 0
        right = 1
        
        for i in range(1, n + 1):
            
            while cur_sum > s and left < i - 1:
                cur_sum -= arr[left]
                left += 1
            
            if cur_sum == s:
                return [left + 1, i]
                
            if i < n:
                cur_sum += arr[i]
        
        return [-1]
