a = [2, 3, 5, 7]
b = a
b[0] = '>'
print(a, b)  # ['>', 3, 5, 7] ['>', 3, 5, 7]

a.append(666)

print(a, b)  # ['>', 3, 5, 7, 666] ['>', 3, 5, 7, 666]
