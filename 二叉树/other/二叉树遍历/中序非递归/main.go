package main

import (
	"GoLeetCode/DataStructure"
	"GoLeetCode/Tools"
	"fmt"
)

/**
前序遍历：先访问根节点，再前序遍历左子树，再前序遍历右子树
-> 中序遍历：先中序遍历左子树，再访问根节点，再中序遍历右子树
后序遍历：先后序遍历左子树，再后序遍历右子树，再访问根节点
*/

func inorderTraversal(root *DataStructure.TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*DataStructure.TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			// 根节点在左节点之后入栈
			stack = append(stack, root)
			root = root.Left
		}

		// 出栈
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}

func main() {
	// [3 5 2 4 1 6]
	fmt.Println(inorderTraversal(Tools.GenerateBinaryTree()))
}