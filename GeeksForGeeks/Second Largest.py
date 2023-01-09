class Solution:
	def print2largest(self,arr, n):
	    max = arr[0]
	    second_max = -1
	    
	    for element in arr:
	        if element > max:
	            save = max
	            max = element
	            second_max = save
	        elif element < max and element > second_max:
	            second_max = element
	    
	    return second_max
