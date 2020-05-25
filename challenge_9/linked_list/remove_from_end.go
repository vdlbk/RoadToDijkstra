package linked_list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 || head == nil {
		return head
	}

	return removeNth(*head, n)
}

func removeNth(head ListNode, n int) *ListNode {
	var ints []int
	for head.Next != nil {
		ints = append(ints, head.Val)
		head = *head.Next
	}
	ints = append(ints, head.Val)

	if n > len(ints) {
		return nil
	}

	idx := len(ints) - n
	ints = append(ints[:idx], ints[idx+1:]...)

	// Remove nth
	return GenerateListFromSlice(ints)
}

// way more opti
func removeNthFromEndAlternative(head *ListNode, n int) *ListNode {

	var fast *ListNode = nil
	var slow *ListNode = nil

	// move fast pointer n steps forward
	for n > 0 {
		if fast == nil {
			fast = head
		} else {
			fast = fast.Next
		}
		n--
	}

	for fast.Next != nil {
		fast = fast.Next

		if slow == nil {
			slow = head
		} else {
			slow = slow.Next
		}
	}

	if slow == nil {
		head = head.Next
	} else {
		//fmt.Printf("slow=%d;fast=%d\n", slow.Val, fast.Val)
		temp := slow.Next.Next
		slow.Next = temp
		temp = nil
	}

	return head
}
