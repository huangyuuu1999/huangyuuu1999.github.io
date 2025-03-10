import bisect


class Solution:
    def hasSameDigits(self, s: str) -> bool:
        while len(s) > 2:
            tmp = []
            for i in range(len(s)-1):
                a, b = int(s[i]), int(s[i + 1])
                c = (a + b) % 10
                tmp.append(str(c))
                s = ''.join(tmp)
        return tmp[0] == tmp[1]

ans = Solution().hasSameDigits("3902")
print(ans)

bisect.bisect_left()