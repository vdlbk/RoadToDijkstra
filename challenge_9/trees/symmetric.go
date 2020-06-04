package trees

func isSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	return isMirror(root, root)
}

func isMirror(root, other *TreeNode) bool {
	if root == nil || other == nil {
		return root == nil && other == nil
	}

	return root.Val == other.Val && isMirror(root.Left, other.Right) && isMirror(root.Right, other.Left)
}
