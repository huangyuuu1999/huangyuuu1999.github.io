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