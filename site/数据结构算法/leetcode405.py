from typing import List


class Solution:
    def minimumCost(self, target: str, words: List[str], costs: List[int]) -> int:
        path = []
        min_cost = float('inf')
        cur_cost = 0
        def dfs():
            nonlocal min_cost, cur_cost
            path_str = ''.join(path)
            if path_str == target:
                min_cost = min(cur_cost, min_cost)
                return
            for word_index in range(len(words)):
                path.append(words[word_index])
                path_str = ''.join(path)
                if path_str == target[:len(path_str)]:
                    cur_cost += costs[word_index]
                    dfs()
                    cur_cost -= costs[word_index]
                    path.pop()
                else:
                    path.pop()
        dfs()
        if min_cost ==float('inf'):
            return -1
        return min_cost

target = "abcdef"
words = ["abdef","abc","d","def","ef"]
costs = [100,1,10,2,5]
ans = Solution().minimumCost(target, words, costs)
print(ans)