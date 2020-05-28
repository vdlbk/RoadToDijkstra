package linked_list

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	m := make(map[*ListNode]bool)

	for {
		if m[head] == true {
			return true
		}
		if head != nil {
			m[head] = true
			head = head.Next
		} else {
			return false
		}
	}

	return false
}

func hasCycleOpti(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
