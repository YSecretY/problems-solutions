class Solution:
    def sum(array):
        sum = 0
        for i in range(0, len(array)):
            sum += array[i]
        return sum
    def MissingNumber(self,array,n):
        ideal_array = []
        for i in range(1, n + 1):
            ideal_array.append(i)
        return sum(ideal_array) - sum(array)
