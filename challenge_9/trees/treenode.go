package trees

import "fmt"

var null = -1 << 31

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t TreeNode) String() string {
	return fmt.Sprintf("Val: %v\n\tLeft: %+v, \n\tRight: %+v", t.Val, t.Left, t.Right)
}

func Insert(ints []int, root *TreeNode, i int) *TreeNode {
	if i < len(ints) && ints[i] != null {

		temp := &TreeNode{Val: ints[i]}
		root = temp
		root.Left = Insert(ints, root.Left, 2*i+1)
		root.Right = Insert(ints, root.Right, 2*i+2)
	}
	return root
}

func GenerateTreeNode(ints []int) *TreeNode {
	if len(ints) == 0 {
		return nil
	}
	return Insert(ints, &TreeNode{}, 0)
}
