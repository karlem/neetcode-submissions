/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    out := []int{}

	var rec func(root *TreeNode)
	rec = func(root *TreeNode) {
		if root == nil {
			return
		}

		rec(root.Left)
		rec(root.Right)

		out = append(out, root.Val)
	}

	rec(root)

	return out
}
