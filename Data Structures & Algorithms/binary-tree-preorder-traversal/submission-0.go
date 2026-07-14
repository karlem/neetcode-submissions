/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
	out := []int{}
    var rec func(node *TreeNode)
	rec = func(node *TreeNode) {
		if node == nil {
			return
		}

		out = append(out, node.Val)

		rec(node.Left)
		rec(node.Right)
	}
	rec(root)

	return out
}
