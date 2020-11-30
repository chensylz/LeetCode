package palindrome_linked_list

/**
234. 回文链表
https://leetcode-cn.com/problems/palindrome-linked-list/
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 翻转后面
	tail := reverse(slow.Next)
	slow.Next = nil

	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}



