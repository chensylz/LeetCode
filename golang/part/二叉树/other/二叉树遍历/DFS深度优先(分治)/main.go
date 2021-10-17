package main

import (
	"GoLeetCode/DataStructure"
	"GoLeetCode/Tools"
	"fmt"
)

// DFS 深度搜索（从上到下） 和分治法区别：
//前者一般将最终结果通过指针参数传入，
//后者一般递归返回结果最后合并

func preorderTranversal(root *DataStructure.TreeNode) []int {
	return divideAndConquer(root)
}

func divideAndConquer(root *DataStructure.TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return nil
	}
	// 分治
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	// 合并
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func main() {
	// [1 2 3 5 4 6]
	fmt.Println(preorderTranversal(Tools.GenerateBinaryTree()))
}