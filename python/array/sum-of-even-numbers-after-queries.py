# coding=utf8

"""
985. 查询后的偶数和
https://leetcode-cn.com/problems/sum-of-even-numbers-after-queries/
"""
from typing import List


class Solution:
    def sumEvenAfterQueries(self, A: List[int], queries: List[List[int]]) -> List[int]:
        answer = []
        if len(queries) == 0:
            return answer

        sum_of_even = sum(x for x in A if x % 2 == 0)

        for query in queries:
            value = query[0]
            index = query[1]
            if A[index] % 2 == 0:
                sum_of_even -= A[index]
            A[index] += value
            if A[index] % 2 == 0:
                sum_of_even += A[index]
            answer.append(sum_of_even)
        return answer


if __name__ == '__main__':
    print(Solution().sumEvenAfterQueries([1,2,3,4], [[1,0],[-3,1],[-4,0],[2,3]]))
