package _76_middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	low, fast := head, head
	for fast.Next != nil {
		low = low.Next
		if fast.Next.Next != nil {
			fast = fast.Next.Next
		} else if fast.Next != nil {
			fast = fast.Next
		}
	}
	return low
}
