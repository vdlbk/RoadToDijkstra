package linked_list

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l := &ListNode{
		Val:  head.Val,
		Next: nil,
	}
	for head.Next != nil {
		l = &ListNode{
			Val:  head.Next.Val,
			Next: l,
		}
		head = head.Next
	}
	return l
}

// Alternative (faster)
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
