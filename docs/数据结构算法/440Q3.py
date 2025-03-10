from bisect import bisect_left
from typing import List


class Solution:
    def numOfUnplacedFruits(self, fruits: List[int], baskets: List[int]) -> int:
        n = len(fruits)
        used = [False] * n
        ans = 0
        for f in fruits:
            index = -1
            while True:
                index = bisect_left(baskets, f, lo=index+1)
                if index >= n:
                    break
                if not used[index]:
                    used[index] = True
                    ans += 1
                    break
        return ans


fs, bs = [7, 5], [8, 5]
Solution().numOfUnplacedFruits(fs, bs)