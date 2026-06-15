/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
	if root == nil {
		return res
	}
	
	q := []*TreeNode{root}

	for len(q) > 0 {
		l := len(q)

		var level []int
		for i := 0; i < l; i++ {
			pop := q[0]
			q = q[1:]

			level = append(level, pop.Val)
			if pop.Left != nil {
				q = append(q, pop.Left)
			}
			if pop.Right != nil {
				q = append(q, pop.Right)
			}
		}

		res = append(res, level)
	}

	return res
}
