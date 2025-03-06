from collections import Counter, deque
import heapq
from typing import List

s = """
a word are a world and a are a word and word word word word
"""

c = Counter(s.split())
print(c)


class TreeNode:
    def __init__(self, val, freq, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
        self.freq = freq

    def __lt__(self, other: "TreeNode"):
        return self.freq < other.freq

    def __repr__(self):
        return f"<{self.val},{self.freq}>"


path = []
str_code_dictionary = {}
code_str_dictionary = {}


def dfs(root: TreeNode):
    if root is None:
        return
    if root.left is root.right:
        code = "".join(path)
        str_code_dictionary[root.val] = code
        code_str_dictionary[code] = root.val

    path.append("0")
    dfs(root.left)
    path.pop()

    path.append("1")
    dfs(root.right)
    path.pop()


nodes = [TreeNode(k, v) for k, v in c.items()]


def build_Tree(node_list: List[TreeNode]) -> TreeNode:
    heapq.heapify(node_list)
    if len(node_list) == 1:
        return node_list[0]
    # 弹出最小的两个节点
    n1, n2 = heapq.heappop(node_list), heapq.heappop(node_list)
    if n1.freq > n2.freq:  # 保证左边的小
        n1, n2 = n2, n1
    parent = TreeNode("○", n1.freq + n2.freq, n1, n2)
    heapq.heappush(node_list, parent)
    return build_Tree(node_list)


def print_tree_graphically(root: TreeNode):
    if not root:
        print("nil")
        return

    # 计算树的深度
    def get_depth(node):
        if not node:
            return 0
        return max(get_depth(node.left), get_depth(node.right)) + 1

    depth = get_depth(root)
    level = 0
    queue = deque([(root, level, 2 ** (depth - 1) - 1)])  # (node, level, position)
    lines = [[" " for _ in range(2**depth - 1)] for _ in range(depth * 2 - 1)]

    while queue:
        node, level, pos = queue.popleft()
        # 打印当前节点
        lines[level * 2][pos] = f"{node.val} ({node.freq})"
        # 打印连接线
        if node.left:
            lines[level * 2 + 1][pos - 1] = "/"
            queue.append((node.left, level + 1, pos - 2 ** (depth - level - 2)))
        if node.right:
            lines[level * 2 + 1][pos + 1] = "\\"
            queue.append((node.right, level + 1, pos + 2 ** (depth - level - 2)))

    for line in lines:
        print("".join(line).center(2**depth * 4))  # 调整居中对齐


root = build_Tree(nodes)
print_tree_graphically(root)

dfs(root)
print(code_str_dictionary)
print(str_code_dictionary)
