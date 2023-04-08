func removeElements(head *ListNode, val int) *ListNode {
    for head != nil && head.Val == val {
        head = head.Next
    }

    if head == nil {
        return head
    }

    prev := head
    for cur := head; cur != nil; cur = cur.Next {
        if cur.Val == val {
            prev.Next = cur.Next
        } else {
            prev = cur
        }
    }

    return head
}
