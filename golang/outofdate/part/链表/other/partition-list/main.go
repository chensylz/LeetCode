package main

import "fmt"

/**
86. 分隔链表
https://leetcode-cn.com/problems/partition-list/
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: 0}
	var saveDummy = dummy
	var saveHead = head

	for head != nil {
		if head.Val < x {
			dummy.Next = &ListNode{Val: head.Val}
			dummy = dummy.Next
		}
		head = head.Next
	}

	head = saveHead

	for head != nil {
		if head.Val >= x {
			dummy.Next = &ListNode{Val: head.Val}
			dummy = dummy.Next
		}
		head = head.Next
	}

	return saveDummy.Next
}

func GenerateLinked(arr []int) *ListNode {
	head := &ListNode{Val: 0}
	var dummy = head
	for _, value := range arr {
		head.Next = &ListNode{Val: value}
		head = head.Next
	}
	return dummy.Next
}

func PrintLinked(root *ListNode) {
	for root != nil {
		fmt.Print(root.Val)
		fmt.Print(" ")
		root = root.Next
	}
	fmt.Println()
}

func main() {
	PrintLinked(partition(GenerateLinked([]int{1,2,4,3,5,2}), 3))
}