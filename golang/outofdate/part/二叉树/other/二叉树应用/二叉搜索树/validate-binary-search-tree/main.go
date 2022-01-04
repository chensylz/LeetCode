package validate_binary_search_tree

/**
98. 验证二叉搜索树
https://leetcode-cn.com/problems/validate-binary-search-tree/
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// ---------------------------中序遍历-----------------------------
// 检查结果列表是否已经有序
func isValidBSTV1(root *TreeNode) bool {
	result := make([]int, 0)
	inOrder(root, &result)
	for i := 0; i < len(result) - 1; i++{
		if result[i] >= result[i + 1] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, result)
	*result = append(*result, root.Val)
	inOrder(root.Right, result)
}

// ---------------------------分治-------------------------------
// 判断左 MAX < 根 < 右 MIN

func isValidBSTV2(root *TreeNode) bool {
	result := helper(root)
	return result.IsValid
}

type ResultType struct {
	IsValid bool
	Max *TreeNode
	Min *TreeNode
}

func helper(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{IsValid: true}
	}
	result := ResultType{}
	left := helper(root.Left)
	right := helper(root.Right)

	if !left.IsValid || !right.IsValid {
		result.IsValid = false
		return result
	}

	if left.Max != nil && left.Max.Val >= root.Val {
		result.IsValid = false
		return result
	}

	if right.Min != nil && right.Min.Val <= root.Val {
		result.IsValid = false
		return result
	}

	result.IsValid = true
	result.Min = root
	if left.Min != nil {
		result.Min = left.Min
	}
	result.Max = root
	if right.Max != nil {
		result.Max = right.Max
	}
	return result
}