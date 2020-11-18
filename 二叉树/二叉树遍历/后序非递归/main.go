package main

/**
前序遍历：先访问根节点，再前序遍历左子树，再前序遍历右子树
中序遍历：先中序遍历左子树，再访问根节点，再中序遍历右子树
-> 后序遍历：先后序遍历左子树，再后序遍历右子树，再访问根节点
*/

import (
	"GoLeetCode/DataStructure"
	"GoLeetCode/Tools"
	"fmt"
)


func postorderTraversal(root *DataStructure.TreeNode) []int {
	 if root == nil {
	 	return nil
	 }

	 result := make([]int, 0)
	 stack := make([]*DataStructure.TreeNode, 0)
	 var lastVisit *DataStructure.TreeNode

	 for root != nil || len(stack) > 0 {
	 	for root != nil {
	 		stack = append(stack, root)
	 		root = root.Left
		}

		node := stack[len(stack) - 1]

		// 根节点需在右节点弹出之后，再弹出
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack) - 1]
			result = append(result, node.Val)
			lastVisit = node
		} else {
			root = node.Right
		}
	 }
	 return result
}

func main() {
	fmt.Println(postorderTraversal(Tools.GenerateBinaryTree()))
}