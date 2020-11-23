package remove_duplicates_from_sorted_list_ii

/**
82. 删除排序链表中的重复元素 II
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
思路：链表头结点可能被删除，所以用 dummy node 辅助删除
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy

	var rmVal int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			// 判断后面还有没有相等的
			rmVal = head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}
