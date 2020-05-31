package trees

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	return check(root, math.MinInt64, math.MaxInt64)
}

func check(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	return root.Val > min && root.Val < max && check(root.Left, min, root.Val) && check(root.Right, root.Val, max)
}
