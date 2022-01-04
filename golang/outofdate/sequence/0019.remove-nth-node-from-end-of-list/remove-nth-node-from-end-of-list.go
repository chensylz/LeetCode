package _019_remove_nth_node_from_end_of_list

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	slow, fast := dummy, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for ;fast != nil; {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next
}
