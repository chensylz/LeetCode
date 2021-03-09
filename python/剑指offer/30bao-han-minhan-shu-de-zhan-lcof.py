# coding=utf8

"""
剑指 Offer 30. 包含min函数的栈
https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/
"""


class MinStack:
    def __init__(self):
        """
        initialize your data structure here.
        """
        self.help_stack = []
        self.stack = []

    def push(self, x: int) -> None:
        self.stack.append(x)
        if not self.help_stack or self.help_stack[-1] >= x:
            self.help_stack.append(x)

    def pop(self) -> None:
        if self.stack.pop() == self.help_stack[-1]:
            self.help_stack.pop()

    def top(self) -> int:
        return self.stack[-1]

    def min(self) -> int:
        return self.help_stack[-1]
