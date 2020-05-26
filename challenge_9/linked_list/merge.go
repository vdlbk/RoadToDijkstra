package linked_list

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var ints []int
	for {
		if l1 == nil {
			for l2.Next != nil {
				ints = append(ints, l2.Val)
				l2 = l2.Next
			}
			ints = append(ints, l2.Val)
			break
		} else if l2 == nil {
			for l1.Next != nil {
				ints = append(ints, l1.Val)
				l1 = l1.Next
			}
			ints = append(ints, l1.Val)
			break
		}

		if l1.Val < l2.Val {
			ints = append(ints, l1.Val)
			l1 = l1.Next
		} else {
			ints = append(ints, l2.Val)
			l2 = l2.Next
		}
	}

	return GenerateListFromSlice(ints)
}
