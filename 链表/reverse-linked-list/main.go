package reverse_linked_list

/**
206. 反转链表
https://leetcode-cn.com/problems/reverse-linked-list/
思路：用一个 prev 节点保存向前指针，temp 保存向后的临时指针
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev

		// 移动
		prev = head
		head = temp
	}
	return prev
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	root := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return root
}
