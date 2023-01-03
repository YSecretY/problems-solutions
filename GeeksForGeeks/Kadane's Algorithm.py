class Solution:
    def maxSubArraySum(self,arr,N):
        max_sum = arr[0]
        cur_sum = 0
        
        for number in arr:
            if cur_sum < 0:
                cur_sum = 0
            cur_sum += number
            max_sum = max(max_sum, cur_sum)
        return max_sum
