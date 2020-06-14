package trees

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	return insertToBST(&TreeNode{}, nums)
}

func insertToBST(root *TreeNode, nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root.Val = nums[mid]
	root.Left = insertToBST(&TreeNode{}, nums[:mid])
	root.Right = insertToBST(&TreeNode{}, nums[mid+1:])
	return root
}

func sortedArrayToBSTOpti(nums []int) *TreeNode {
	var root *TreeNode
	if len(nums) > 0 {
		mid := len(nums) / 2
		if mid > 0 {
			root = &TreeNode{
				Val:   nums[mid],
				Left:  sortedArrayToBSTOpti(nums[:mid]),
				Right: sortedArrayToBSTOpti(nums[mid+1:]),
			}
		} else {
			root = &TreeNode{
				Val: nums[mid],
			}
		}
	}

	return root
}
