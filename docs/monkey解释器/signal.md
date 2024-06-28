# 读取字符，判断是不是界限符和运算符

### 读取字符
1. **让词法解析器能够读取字符** 请写一段代码，每次能够读取字符串中的一个字符并打印出来
    ```python
    for c in range(len(s)):
        print(c)
    ```

### 可穷举与不可穷举

### 约定符号集

### 判断符号（界限符和运算符）

2. **让词法解析器能够识别运算符、界限符** 请写一段代码，如果遇到 `=+,;(){}` 这些运算符和界限符，就当做token按顺序保存起来，其他的字符直接丢弃
    ```python
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
    你也可以使用字典，但无论如何这段代码只能识别可穷举的界限符和运算符，针对关键字和标识符无能为力。