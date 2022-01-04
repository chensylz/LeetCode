package main

import (
	"GoLeetCode/outofdate/DataStructure"
	"GoLeetCode/outofdate/Tools"
	"fmt"
)

func levelOrder(root *DataStructure.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*DataStructure.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)

		l := len(queue)
		for i := 0; i < l; i++ {
			// 出
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			} else if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, list)
	}

	return result
}

func main() {
	fmt.Println(levelOrder(Tools.GenerateBinaryTree()))
}
