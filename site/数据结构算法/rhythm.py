rhythm = [
    ['一', '三', '七', '八'],
    ['零', '十'],
    ['五', '九'],
    ['二', '四', '六']
]

ans = []
path = []

def backtrace(i: int):
    if i == len(rhythm):
        ans.append(''.join(path) + '\n')
        return
    for _, ch in enumerate(rhythm[i]):
        path.append(ch)
        backtrace(i+1)
        path.pop()

backtrace(0)

with open("res.txt", 'w') as f:
    f.writelines(ans)