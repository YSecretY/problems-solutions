class Solution:
	def binarysearch(self, arr, n, k):
      high = n - 1
      low = 0
      while low <= high:
          mid = (low + high) // 2
          if arr[mid] == k:
              return mid
          if arr[mid] > k:
              high = mid - 1
          elif arr[mid] < k:
              low = mid + 1
      return -1
