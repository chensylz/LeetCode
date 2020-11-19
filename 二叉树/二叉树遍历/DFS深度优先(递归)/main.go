package main

import (
	"GoLeetCode/DataStructure"
	"GoLeetCode/Tools"
	"fmt"
)

func preorderTranversal(root *DataStructure.TreeNode) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(root *DataStructure.TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}

func main() {
	// [1 2 3 5 4 6]
	fmt.Println(preorderTranversal(Tools.GenerateBinaryTree()))
}
