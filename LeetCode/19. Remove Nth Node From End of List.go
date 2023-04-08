func removeNthFromEnd(head *ListNode, n int) *ListNode {
    mp := make(map[int]*ListNode)
    var l, cur = 0, head
    for ; cur != nil; l, cur = l+1, cur.Next {
        mp[l+1] = cur
    }

    if l == 1 {
        return nil
    }

    el := mp[l-n+1]
    if el == head {
        return head.Next
    }
    mp[l-n].Next = el.Next
    

    return head
}
