package main

import (
	"GoLeetCode/DataStructure"
	"GoLeetCode/Tools"
	"fmt"
)

/**
-> 前序遍历：先访问根节点，再前序遍历左子树，再前序遍历右子树
中序遍历：先中序遍历左子树，再访问根节点，再中序遍历右子树
后序遍历：先后序遍历左子树，再后序遍历右子树，再访问根节点
 */


func preorderTraversal(root *DataStructure.TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

func main() {
	preorderTraversal(Tools.GenerateBinaryTree())
}