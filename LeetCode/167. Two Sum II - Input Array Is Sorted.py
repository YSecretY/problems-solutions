class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        small = 0
        big = len(numbers) - 1
        while numbers[small] + numbers[big] != target:
            sum = numbers[small] + numbers[big]
            if sum > target:
                big -= 1
            else:
                small += 1
        return [small + 1, big + 1]
