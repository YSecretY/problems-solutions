class Solution:
    def isValid(self, s: str) -> bool:
        mp = {")":"(", "]":"[", "}":"{"}
        stack = []

        for brace in s:
            if brace in mp:
                if stack and stack[-1] == mp[brace]:
                    stack.pop()
                else: 
                    return False
            else:
                stack.append(brace)
        
        return True if not stack else False
