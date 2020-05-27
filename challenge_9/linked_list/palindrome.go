package linked_list

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	reversed := reverse(head)
	for head != nil {
		if head.Val == reversed.Val {
			head = head.Next
			reversed = reversed.Next
		} else {
			return false
		}
	}
	return true
}

func reverse(head *ListNode) *ListNode {
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

func isPalindromeOpti(head *ListNode) bool {
	if head == nil {
		return true
	}
	slice := convertListToSlice(head)
	front := 0
	back := len(slice) - 1
	for front < back {
		if slice[front] != slice[back] {
			return false
		}
		front++
		back--
	}
	return true
}

func convertListToSlice(head *ListNode) []int {
	var slice []int
	for true {
		slice = append(slice, head.Val)
		if head.Next != nil {
			head = head.Next
		} else {
			break
		}
	}
	return slice
}
