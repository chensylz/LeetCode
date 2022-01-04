package Tools

import (
	"GoLeetCode/outofdate/DataStructure"
)

func GenerateBinaryTree() *DataStructure.TreeNode {
	root := generateNode(1)
	root.Left = generateNode(2)
	root.Left.Left = generateNode(3)
	root.Left.Right = generateNode(4)
	root.Left.Left.Right = generateNode(5)
	root.Right = generateNode(6)

	return root
}

func generateNode(val int) *DataStructure.TreeNode {
	return &DataStructure.TreeNode{
		Val: val,
	}
}
