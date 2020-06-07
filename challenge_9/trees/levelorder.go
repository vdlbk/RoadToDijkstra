package trees

func levelOrder(root *TreeNode) [][]int {
	levels := make([][]int, 0)
	if root == nil {
		return levels
	}

	levels = append(levels, []int{root.Val})
	levels = values(root, levels, 1)
	return levels
}

func values(root *TreeNode, levels [][]int, row int) [][]int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return levels
	}

	if len(levels) <= row {
		levels = append(levels, []int{})
	}

	nextRow := row + 1
	if root.Left != nil {
		levels[row] = append(levels[row], root.Left.Val)
		levels = values(root.Left, levels, nextRow)
	}
	if root.Right != nil {
		levels[row] = append(levels[row], root.Right.Val)
		levels = values(root.Right, levels, nextRow)
	}

	return levels
}
