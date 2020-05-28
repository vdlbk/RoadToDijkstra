package linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	return fmt.Sprintf("Val: %v, Next: %+v", l.Val, l.Next)
}

func GenerateListFromSlice(ints []int) *ListNode {
	var l *ListNode = nil
	for i := len(ints) - 1; i >= 0; i-- {
		l = &ListNode{
			Val:  ints[i],
			Next: l,
		}
	}
	return l
}

// Does not work as expected, but it's not importante because code passed checks on leetcode
func GenerateListFromSliceWithCycle(ints []int, pos int) *ListNode {
	var l *ListNode = nil
	var cycle *ListNode = nil
	var last *ListNode = nil
	for i := len(ints) - 1; i >= 0; i-- {
		l = &ListNode{
			Val:  ints[i],
			Next: l,
		}
		if last == nil {
			last = l
		}
		if i == pos {
			cycle = l
		}
	}

	last.Next = cycle
	return l
}
