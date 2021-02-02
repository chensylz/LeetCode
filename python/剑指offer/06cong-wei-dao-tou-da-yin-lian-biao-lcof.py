# coding=utf8

"""
剑指 Offer 06. 从尾到头打印链表
https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
"""
from typing import List


class ListNode:
   def __init__(self, x):
       self.val = x
       self.next = None


class Solution:
    def reversePrint(self, head: ListNode) -> List[int]:
        node_arr = []
        while head is not None:
            node_arr.append(head.val)
            head = head.next
        res = []
        for i in range(len(node_arr) - 1, 0, -1):
            res.append(node_arr[i])
        return res
