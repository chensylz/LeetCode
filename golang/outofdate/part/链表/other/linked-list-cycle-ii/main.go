package linked_list_cycle_ii

/**
142. 环形链表 II
https://leetcode-cn.com/problems/linked-list-cycle-ii/

思路： 快慢指针， 设头到入口点为a, 入口点到相遇点为b，相遇点到入口点为c
当两个指针相遇时， a + n(b+c) + b = 2(a + b)
化简： a = (n - 1)(b + c) + c
n >= 1, 所以，当n=1时，a = c，即，从相遇点到入口点的距离为头部到入口点的距离
 */
type ListNode struct {
	Val int
	Next *ListNode
}


func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast := head.Next
	slow := head

	for fast != nil && fast.Next != nil {
		if fast == slow {
			// 若两个相遇了，那么代表此时需要算出a的长度，再来一个指针算长度
			tempHead := head
			slow = slow.Next // 注意
			for tempHead != slow {
				tempHead = tempHead.Next
				slow = slow.Next
			}
			return tempHead
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}

