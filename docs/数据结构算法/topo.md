# 拓扑排序

## 图的表示

## 拓扑排序算法

### python graphlib库
```python
from graphlib import TopologicalSorter

g = {
    0: [1],
    1: [2],
    2: [3, 4, 5],
    4: [5]
}

ts = TopologicalSorter(g)
res = ts.static_order()

print(list(res))  # [3, 5, 4, 2, 1, 0]
```
### BFS方法

### DFS方法
使用states 标记所有节点的访问状态
## Leetcode 例题

### 课程表2

### 课程表1