package binary_tree_level_order_traversal_ii

/**
107. 二叉树的层次遍历 II
https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := make([][]int, 0)

	for len(queue) > 0 {
		list := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, list)
	}
	reverse(result)
	return result
}

func reverse(nums[][]int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}