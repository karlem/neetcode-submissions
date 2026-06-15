/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	
	var rec func(*TreeNode)
	rec = func(curr *TreeNode) {
		if curr == nil {
			return
		}

		rec(curr.Left)
		res = append(res, curr.Val)
		rec(curr.Right)
	}
	rec(root)
	return res
}
