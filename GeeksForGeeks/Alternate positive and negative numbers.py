class Solution:
    def rearrange(self,arr, n):
        pos_arr = []
        neg_arr = []
        
        for element in arr:
            if element >= 0:
                pos_arr.append(element)
            else:
                neg_arr.append(element)
        if len(pos_arr) == 0 or len(neg_arr) == 0:
            return arr
        
        min_arr_size = min(len(pos_arr), len(neg_arr))
        
        j = 0
        k = 0
        while k < min_arr_size and j < n:
            arr[j] = pos_arr[k]
            arr[j + 1] = neg_arr[k]
            j += 2
            k += 1
        
        if len(pos_arr) > len(neg_arr):
            while k < len(pos_arr) and j < n:
                arr[j] = pos_arr[k]
                k += 1
                j += 1
        else:
            while k < len(neg_arr) and j < n:
                arr[j] = neg_arr[k]
                k += 1
                j += 1
        
        return arr
