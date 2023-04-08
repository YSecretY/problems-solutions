func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil  {
        return head
    }

    var prev, cur *ListNode = nil, head

    for cur != nil {
        temp := cur.Next
        cur.Next = prev
        prev = cur
        cur = temp
    }

    return prev
}
