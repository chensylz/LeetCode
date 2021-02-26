# coding=utf8

"""
剑指 Offer 18. 删除链表的节点
https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
"""

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def deleteNode(self, head: ListNode, val: int) -> ListNode:
        if not head:
            return head

        tmp_head = head
        if tmp_head.val == val:
            return tmp_head.next

        while tmp_head and tmp_head.next:
            if tmp_head.next.val == val:
                tmp_head.next = tmp_head.next.next
            tmp_head = tmp_head.next

        return head
