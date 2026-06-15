/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return valid(root, math.MinInt, math.MaxInt)
}


func valid(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}

	val := root.Val
	if val <= left || val >= right {
		return false
	}

	return valid(root.Left,left, val) && valid(root.Right, val, right)
}