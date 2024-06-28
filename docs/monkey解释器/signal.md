# 读取字符，判断是不是界限符和运算符

### 读取字符
1. **让词法解析器能够读取字符** 请写一段代码，每次能够读取字符串中的一个字符并打印出来
    ```python
    s = "let x = 5; x = x + 7;"
    for c in range(len(s)):
        print(c)
    ```
    这样写可以实现遍历远吗字符串的功能，但是想一下，有时候我们需要识别的不是单个字符，而是一个单词。这样按照字符串的长度来遍历，就不是一个明智的做法了。
    ```python
    # 简单的实现 不用oop
    s = "let x = 5; x = x + 7;"
    cur_pos, next_pos = -1, 0
    cur_ch = ''
    read_ch = []
    while next_pos < len(s):
        cur_ch = s[next_pos]
        read_ch.append(cur_ch)
        cur_pos = next_pos
        next_pos += 1
    print(cur_ch)
    ```
### 可穷举与不可穷举

运算符和界限符，在一门特定的语言中，数量是有限的，在上面的读取字符的代码中，我们可以加一些逻辑：每读一个就判断一下是不是某个界限符或者运算符，这样就可以识别出所有的运算符和界限符。
```python
# 简单的实现 不用oop
s = "let x = 5; x = x + 7;"
cur_pos, next_pos = -1, 0
cur_ch = ''
read_ch = []
while next_pos < len(s):
    cur_ch = s[next_pos]
    if cur_ch in ['=', '+', '(', ')', '{', '}']:  # 可以拓展这个列表，完全取决于如何约定
        read_ch.append(cur_ch)
    cur_pos = next_pos
    next_pos += 1
print(read_ch)
```
请注意我们可以逐一判断的理由是：界限符和运算符本身是有限的。相反，对于标识符例如变量名，他们是没有数量限制或者数量太多的，你可以起任意合法的变量名，我们不能写一堆case来判断这个字母是不是`a` 是不是`b`...

增加了判断的逻辑之后，我们得词法分析器可以判断界限符和运算符了，但是没有判断其他类型例如标识符，关键字和字面量的能力。


### 判断符号（界限符和运算符）

2. **让词法解析器能够识别运算符、界限符** 请写一段代码，如果遇到 `=+,;(){}` 这些运算符和界限符，就当做token按顺序保存起来，其他的字符直接丢弃
    有了上面的代码，我们得第二个编程练习实际上已经被回答了，我们只需要对这些可穷举的运算符和界限符逐一判断即可。
    ```python
    # 使用简单的for遍历字符串
    tokens = []
    s = "{x =x+1}"
    for c in range(len(s)):
        if c == '=':
            tokens.append('=')
        elif c == ';':
            tokens.append(';')
        elif c == '{':
            tokens.append('{')
        ...  # 我们要把所有需要识别的字符都列出来（可以使用switch case简化）
        else:
            print('other, drop.')
    print(tokens)
    ```
    你也可以使用字典，但无论如何这段代码只能识别可穷举的界限符和运算符，针对关键字和标识符无能为力；不过在下一节对代码稍加改进，我们将可以识别关键字和标识符。