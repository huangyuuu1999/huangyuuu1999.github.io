def partition(arr, p, r):
    assert p <= r
    pivot = arr[0]
    left_part, right_part = [], []  # 保存左边和右边
    for n in arr:
        if n < pivot:
            left_part.append(n)
        elif n > pivot:
            right_part.append(n)
    new_arr = left_part + [pivot] + right_part  # 更新原来的数组
    for i in range(len(arr)):
        arr[i] = new_arr[i]
    return new_arr


arr = [5, 8, 2, 1, 9, 3, 6, 4]
print("before partition:", arr)
new_arr = partition(arr, 0, len(arr) - 1)
print("after  partition:", arr)

print("returned new arr:", new_arr)

def lomuto_oartition(arr, p, r):
    pivot = arr[p]
    j = p
    for i in range(p+1, r+1):
        if arr[i] < pivot:
            j += 1
            arr[i], arr[j] = arr[j], arr[i]
    arr[p], arr[j] = arr[j], arr[p]
    return j
