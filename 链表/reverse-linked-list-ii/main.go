package reverse_linked_list_ii

/**
92. 反转链表 II
https://leetcode-cn.com/problems/reverse-linked-list-ii/
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}
	// 思路：先遍历到m处，翻转，再拼接后续，注意指针处理
	// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
	var prev *ListNode
	// 头部变化所以使用dummy node
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	i := 0
	// 最开始：0->1->2->3->4->5->nil
	// 先遍历到M处
	for i < m {
		prev = head
		head = head.Next
		i++
	}
	// 遍历之后： 1(pre)->2(head, start)->3->4->5->NULL
	// i = 1
	var start *ListNode
	var end = head
	for head != nil && i <= n {
		// 第一次循环： 1(prev) nil<-2(end, start) 3(head)->4->5->nil
		// 第二次循环： 1(prev) nil<-2(end)<-3(start) 4(head)->5->nil
		// 第三次循环： 1(prev) nil<-2(end)<-3<-4(start) 5(head)->nil
		temp := head.Next
		head.Next = start
		start = head
		head = temp
		i++
	}
	prev.Next = start
	end.Next = head

	return dummy.Next
}
