package _876_middle_of_the_linked_list

type ListNode struct {
	Val int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	n := 1
	tempPtr := head
	for tempPtr.Next != nil {
		n ++
		tempPtr = tempPtr.Next
	}
	for i := 0; i < n / 2; i++ {
		head = head.Next
	}
	return head
}

func middleNode2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	return slow
}
