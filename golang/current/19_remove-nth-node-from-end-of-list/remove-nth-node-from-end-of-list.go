package _9_remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := &ListNode{Next: head}
	low, fast := tmp, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		low = low.Next
		fast = fast.Next
	}
	low.Next = low.Next.Next
	return tmp.Next
}
