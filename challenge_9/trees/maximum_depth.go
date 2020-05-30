package trees

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	depth := 1
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		depth += left
	} else {
		depth += right
	}

	return depth
}

// Best answer on leetcode but benchmarks show small differences
func maxDepthOpti(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Max(maxDepth(root.Right), maxDepth(root.Left))
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
