package reorder_list

/**
143. 重排链表
https://leetcode-cn.com/problems/reorder-list/
 */
// 思路：找到中点断开，翻转后面部分，然后合并前后两个链表

type ListNode struct {
	Val int
	Next *ListNode
}

func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}
	middle := findMiddle(head)
	tail := reverseList(middle.Next) // 细节 middle.Next
	middle.Next = nil
	head = mergeTowList(head, tail)
}

func findMiddle(head *ListNode) *ListNode {
	fast := head.Next
	low := head
	for fast != nil && fast.Next != nil { // 细节fast.next != nil
		fast = fast.Next.Next
		low = low.Next
	}
	return low
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		// nil<-1(pre) 2->3->4
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

func mergeTowList(left, right *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	head := dummy
	isLeft := true
	for left != nil && right != nil {
		if isLeft {
			head.Next = left
			left = left.Next
		} else {
			head.Next = right
			right = right.Next
		}
		head = head.Next
		isLeft = !isLeft
	}

	if left != nil {
		head.Next = left
	}

	if right != nil {
		head.Next = right
	}

	return dummy.Next
}