ListNode* deleteDuplicates(ListNode* head) {
    if (!head)
        return nullptr;
    ListNode* save = head;
    while (head -> next) {
        if (head -> next -> val == head -> val)
            head -> next = head -> next -> next;
        else head = head -> next;
    }
    return save;
}
