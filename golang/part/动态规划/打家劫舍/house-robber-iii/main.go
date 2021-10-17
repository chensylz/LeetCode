package house_robber_iii

/**
337. 打家劫舍 III
https://leetcode-cn.com/problems/house-robber-iii/
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	memo = make(map[*TreeNode]int)
)

func robV1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if value, ok := memo[root]; ok {
		return value
	}
	doIt := root.Val
	if root.Left != nil {
		doIt += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		doIt += rob(root.Right.Left) + rob(root.Right.Right)
	}
	notDoIt := rob(root.Left) + rob(root.Right)
	result := max(doIt, notDoIt)
	memo[root] = result

	return result
}

func rob(root *TreeNode) int {
	res := dp(root)
	return max(res[0], res[1])
}
/* 返回⼀个⼤⼩为 2 的数组 arr
arr[0] 表⽰不抢 root 的话，得到的最⼤钱数
arr[1] 表⽰抢 root 的话，得到的最⼤钱数 */
func dp(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := dp(root.Left)
	right := dp(root.Right)

	// 抢
	doIt := root.Val + left[0] + right[0]
	// 不抢，可以下家也不抢
	notDoIt := max(left[0], left[1]) + max(right[0], right[1])

	return []int{notDoIt, doIt}
}


func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}