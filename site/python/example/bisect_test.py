from bisect import bisect_left


a = [2, 3, 5, 7, 9, 10]
pos = bisect_left(a, 8)
print(pos)

nums = [1,3,5,6]
target = 5
ans = bisect_left(nums, target)
print(ans)