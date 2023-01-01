class Solution:
    def findIntersection(self, head1, head2):
        hashset = set()
        while head2:
            hashset.add(head2.data)
            head2 = head2.next
        
        newHead = None
        newTail = None
        
        while head1:
            if head1.data in hashset:
                if newHead is None:
                    newHead = newTail = Node(head1.data)
                else:
                    newTail.next = Node(head1.data)
                    newTail = newTail.next
            head1 = head1.next
            
        return newHead
