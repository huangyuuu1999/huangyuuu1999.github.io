import heapq
from typing import List

class Solution:
    def maxSum(self, grid: List[List[int]], limits: List[int], k: int) -> int:
        m, n = len(grid), len(grid[0])
        choices = []
        for i in range(m):
            cpy = [-x for x in grid[i]]
            heapq.heapify(cpy)
            print(cpy)
            for c in range(min(limits[i], n)):
                choices.append(heapq.heappop(cpy))
        heapq.heapify(choices)
        ans = 0
        for _ in range(k):
            ans += heapq.heappop(choices)
        return -ans

grid = [[1,2],[3,4]]
limits = [1,2]
k = 2
ans = Solution().maxSum(grid, limits, k)
print(ans)

grid2 = [[5,3,7],[8,2,6]] 
limits2 = [2,2]
k2 = 3
ans2 = Solution().maxSum(grid2, limits2, k2)
print(ans2)
