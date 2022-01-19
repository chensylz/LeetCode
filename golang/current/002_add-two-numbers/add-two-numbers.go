package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	value1, value2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			value1 = 0
		} else {
			value1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			value2 = 0
		} else {
			value2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (value1 + value2 + carry) % 10}
		carry = (value1 + value2 + carry) / 10
		current = current.Next
	}
	return head.Next
}
