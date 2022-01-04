package sort_list

/**
148. 排序链表
https://leetcode-cn.com/problems/sort-list/
 */

type ListNode struct {
	Val int
	Next *ListNode
}


func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	// 如果只有一个节点
	if head == nil || head.Next == nil {
		return head
	}

	// 找到中点
	middle := findMiddle(head)

	// 断开中点
	tail := middle.Next
	middle.Next = nil

	left := mergeSort(head)
	right := mergeSort(tail)

	result := mergeTowList(left, right)

	return result
}

func findMiddle(head *ListNode) *ListNode {
	// 快慢指针
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func mergeTowList(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	head := dummy
	for left != nil && right != nil {
		if left.Val > right.Val {
			head.Next = right
			right = right.Next
		} else {
			head.Next = left
			left = left.Next
		}
		head = head.Next
	}

	if left != nil {
		head.Next = left
	}

	if right != nil {
		head.Next = right
	}

	return dummy.Next
}