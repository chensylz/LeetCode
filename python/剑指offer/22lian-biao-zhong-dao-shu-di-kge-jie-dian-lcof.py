# coding=utf8

"""
剑指 Offer 22. 链表中倒数第k个节点
https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
"""


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def getKthFromEnd(self, head: ListNode, k: int) -> ListNode:
        low, fast = head, head

        for _ in range(k):
            if not fast:
                return
            fast = fast.next
        while fast:
            fast, low = fast.next, low.next
        return low