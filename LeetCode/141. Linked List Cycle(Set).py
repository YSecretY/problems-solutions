class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        hashset = set()

        while head:
            if head in hashset:
                return True
            hashset.add(head)
            head = head.next
        
        return False
