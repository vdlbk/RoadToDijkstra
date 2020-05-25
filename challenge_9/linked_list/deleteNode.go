package linked_list

func deleteNode(node *ListNode, x int) {
	if node == nil || node.Next == nil {
		return
	}
	if node.Val == x {
		*node = *node.Next
	} else {
		deleteNode(node.Next, x)
	}
}

// WTF
// the function signature is actually compose with a single param which contain the node to remove. So it looks like a dumb answer to me
func deleteNodeAnswer(node *ListNode) {
	*node = *node.Next
}
