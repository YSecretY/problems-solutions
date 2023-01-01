class Solution:
  def binarysearch_helper(self, arr, n, k, low, high):
      if high < low:
          return -1
      mid = (high + low) // 2
      if arr[mid] == k:
          return mid
      if arr[mid] < k:
          return self.binarysearch_helper(arr, n, k, mid + 1, high)
      if arr[mid] > k:
          return self.binarysearch_helper(arr, n, k, low, mid - 1)
	def binarysearch(self, arr, n, k):
      if n == 0:
          return -1
      else: return self.binarysearch_helper(arr, n, k, 0, n - 1)
