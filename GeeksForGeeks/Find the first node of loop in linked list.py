""" Node Class
    class Node:
        def __init__(self, data):   # data -> value stored in node
            self.data = data
            self.next = None
"""
class Solution:
    def findFirstNode(self, head):
        
        while (head.next):
            head.data = -head.data
            head = head.next
            if head.data < 0:
                return -head.data
        
        return -1
