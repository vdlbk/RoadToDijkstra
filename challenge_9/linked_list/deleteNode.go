package linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func generateListFromSlice(ints []int) *ListNode {
	var l *ListNode = nil
	for i := len(ints) - 1; i >= 0; i-- {
		l = &ListNode{
			Val:  ints[i],
			Next: l,
		}
	}
	return l
}

func (l ListNode) String() string {
	return fmt.Sprintf("Val: %v, Next: %+v", l.Val, l.Next)
}
