class Solution:
    def search(self, arr : list, low : int, high : int, target : int):
        while low <= high:
            
            mid = (low + high) // 2
            if arr[mid] == target:
                return mid
            
            # left portion
            if arr[low] <= arr[mid]:
                if target > arr[mid] or target < arr[low]:
                    low = mid + 1
                else:
                    high = mid - 1
            # right portion
            else:
                if target < arr[mid] or target > arr[high]:
                    high = mid - 1
                else:
                    low = mid + 1
                    
        return -1
