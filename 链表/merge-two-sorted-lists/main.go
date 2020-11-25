package merge_two_sorted_lists

/**
21. 合并两个有序链表
https://leetcode-cn.com/problems/merge-two-sorted-lists/
 */
type ListNode struct {
	Val int
	Next *ListNode
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := &ListNode{Val: 0}
	var l3 = dummy

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			l3.Next = l2
			l2 = l2.Next
		} else {
			l3.Next = l1
			l1 = l1.Next
		}
		l3 = l3.Next
	}

	if l1 != nil {
		l3.Next = l1
	}
	if l2 != nil {
		l3.Next = l2
	}
	return dummy.Next
}