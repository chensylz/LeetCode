# coding=utf8

"""
剑指 Offer 09. 用两个栈实现队列
https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
"""

class CQueue:

    def __init__(self):
        self.stack01 = []
        self.stack02 = []

    def appendTail(self, value: int) -> None:
        n = len(self.stack02)
        for i in range(n):
            self.stack01.append(self.stack02.pop())
        self.stack01.append(value)

    def deleteHead(self) -> int:
        n = len(self.stack01)
        for i in range(n):
            self.stack02.append(self.stack01.pop())
        if len(self.stack02) == 0:
            return -1
        return self.stack02.pop()


if __name__ == '__main__':
    c = CQueue()
    print(c.deleteHead())
    print(c.appendTail(5))
    print(c.appendTail(2))
    print(c.deleteHead())
    print(c.deleteHead())
