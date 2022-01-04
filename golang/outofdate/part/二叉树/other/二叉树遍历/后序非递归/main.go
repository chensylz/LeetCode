package main

/**
前序遍历：先访问根节点，再前序遍历左子树，再前序遍历右子树
中序遍历：先中序遍历左子树，再访问根节点，再中序遍历右子树
-> 后序遍历：先后序遍历左子树，再后序遍历右子树，再访问根节点
*/

import (
	"GoLeetCode/outofdate/DataStructure"
	"GoLeetCode/outofdate/Tools"
	"fmt"
)

func postorderTraversal(root *DataStructure.TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*DataStructure.TreeNode, 0)
	var lastVisitNode *DataStructure.TreeNode

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 根节点先出
		node := stack[len(stack)-1]
		if node.Right == nil || lastVisitNode == node.Right {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			lastVisitNode = node
		} else {
			root = node.Right
		}
	}

	return result
}

func main() {
	// [5 3 4 2 6 1]
	fmt.Println(postorderTraversal(Tools.GenerateBinaryTree()))
}
