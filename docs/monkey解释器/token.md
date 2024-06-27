# 词法解析

无论是编译型还是解释型语言，为了能够把编程语言转化为计算机能理解的内容，首先就要对用户输入的内容进行拆分，拆分的结果就是token。
完成这件 **识别、拆分** 任务的代码模块，就是词法解析器，是一种处理文本（符合特定语法规定的）的工具。

## 词法解析器
词法分析器往往是编译或者解释的第一步，他将源代码抽象成为token序列。
词法分析器（lexer）也叫词法单元生成器（tokenizer）或者扫描器（scanner）。

### 词法分析器的功能
词法分析器是一些代码，它们能够实现这样的功能：把用户写的源代码，转换成一堆词法单元（token）。就像是你在翻译文言文的时候，首先需要把一句话拆分成为许多不可再分的单元。

举个例子：你要翻译“庭有枇杷树”，首先就要把这句话拆分为 `["庭","有","枇杷","树"]` 或者 `["庭","有","枇杷树"]`，你要做的事情就是拆分，得到一个单词（有的单词就是一个字）序列，而**不必关心**这句话表达什么意思，是不是符合古汉语的语法。

### 词法解析小案例
对于一门像样的编程语言写的代码，例如下面的这一句：
```js
// 注意这不是任何一种已经存在的语言，
// 只是我自己规定的假想象一种语言的语法，尽管他看上去有许多其他语言的影子
let x = 5 + 5;
```
这门语言的词法解析器应该能辨认出：
最开始有一个关键字let，然后是一个变量的名字x，后面是赋值符号=，加法符号+，还有两个数字5。
词法解析器应该按顺序给我们这样的结果

```python
[
    LET, # 关键字，语言规定的，例如if else while struct await func 等
    IDENTIFIER('x'), # 标识符 变量名x, abc student_age, 函数名 add, push_back
    EQUAL_SIGN, # 符号 + - * / % & > < <= || ? 等
    INTEGER(5), # 数字字面量 
    PLUS_SIGN, 
    INTEGER(5)
]
```

这个结果中的每一个成员，都是一个词法单元token，每一个词法单元，都在源代码中有对应。每一个词法单元token都有自己的类型，主要有这几类
- 关键字 ——— 由一门语言自己规定，不同的语言关键字表不同，例如`c语言`的`include typedef if else while goto` 
- 运算符 加`+`减`-`乘`*`除`/`，取模`%`，与`&`或`|`非`^`关系运算，大于等于`>=`, 有的语言还有三元运算`？`符号，指针取值符`->`，解引用符`*`等
- 界限符 `{} () [] <>` 等, 用于在语句中划分出小的作用域，或者数组根据下标取值等
- 标识符 变量名`x y z abc p_name`，函数名`func add getEventLoop`等
- 字面量 写一个数字 `42, 0, 0xfff, 3.24` 或者一个字符串 `"man out!"`

这些类型的token再加上一些空格换行，注释等，大概就是一门语言中你能写的所有东西。

### 词法分析器：一步步实现
朴素的想法就是把源代码当做字符串去扫描，正如“扫描器”这个名称暗示的那样。
这里我打算用一些小小的编程练习来引入，一步一步地添加功能，直到实现一个简陋的词法解析器。

1. 请写一段代码，每次能够读取字符串中的一个字符并打印出来
    ```python
    for c in range(len(s)):
        print(c)
    ```
2. 请写一段代码，如果遇到 `=+,;(){}` 这些运算符和界限符，就当做token按顺序保存起来，其他的字符直接丢弃
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

3. 完善2，使之能够识别关键字`let`,和标识符`xyz`, `add`, ...
    
    当发现是字母开头时，要往后读取直到遇到非字母，再判断这个连续的字符序列是标识符还是关键字。

4. 完善3，使之能够识别整数字面量

    如果字符是数字字符，就继续读取直到遇到非字符，再把这个数字字符串转换成整数。

### 词法分析器demo
```python
class Token:
    def __init__(self, token_type, token_value=None) -> None:
        self.token_type = token_type
        self.token_value = token_value

    def __str__(self) -> str:
        return f'({self.token_type}\t{self.token_value})'

EOF = 'EOF'

ASSIGN = '='
PLUS = '+'
COMMA = ','
SEMICOLON = ';'
LPARENT = '('
RPARENT = ')'
LBRACE = '{'
RBRACE = '}'

LET = 'LET'
FUNC = 'FUNC'

IDENT = 'IDENT' # 标识符 abc, add, foobar, x, y_1

TOKEN_RULE = {
    # EOF 和 非法字符
    'EOF': EOF,
    # 界限符和运算符
    '+': PLUS,
    '=': ASSIGN,
    ',': COMMA,
    ';': SEMICOLON,
    '(': LPARENT,
    ')': RPARENT,
    '{': LBRACE,
    '}': RBRACE,
    # 关键字
    'let': LET,
    'func': FUNC
}


class Lexer:
    def __init__(self, source_code):
        self.token_list = []
        self.source = source_code
        self.cur_pos = -1
        self.next_pos = 0
        self.cur_ch = ''

    def read_char(self):
        if self.next_pos >= len(self.source):
            self.cur_ch = 'EOF' # 源代码已读取完毕
        else:
            self.cur_ch = self.source[self.next_pos]
        self.cur_pos = self.next_pos
        self.next_pos += 1
        return self.cur_ch

    def look_up_world(self):  # 查看一整个连续字符序列，判断类型
        assert self.cur_ch.isalpha()  # 当lexer.ch 是字母 的时候才会使用这个函数
        word_start = self.cur_pos
        while self.cur_ch.isalpha():  # 往后读，直到遇到非字符
            self.read_char()
        word = self.source[word_start:self.cur_pos]  # 此时cur_pos对应的self.ch已经是非字母，是单词的下一位
        # 让当前读取位置 往前回一位
        self.cur_pos -= 1
        self.next_pos -= 1
        self.ch = self.source[self.cur_pos]
        return word


    def next_token(self) -> Token:  # 目前只能识别运算符号+=和 界限符号,;(){}
        ch = self.read_char()
        if ch in [' ', '\n', '\t', '\r']:
            ch = self.read_char()
        if ch in TOKEN_RULE:
            token = Token(TOKEN_RULE[ch])
            self.token_list.append(token)
        elif ch.isalpha():
            word = self.look_up_world()
            if word in TOKEN_RULE:  # 是关键字
                token = Token(TOKEN_RULE[word])
                self.token_list.append(token)
            else:  # 是标识符
                token = Token(IDENT, word)
                self.token_list.append(token)
        else:
            token = Token('Unknown')
        return token

source = "{(y}let),;x+ afuncbc= func("  # 注意空格也会被识别为unknown
lx = Lexer(source)
for _ in range(len(source)):
    token = lx.next_token()  # 思考为什么读取 len(source)次，最后会出现很多EOF？（提示：token的数量和字符的数量关系）
    print(token.token_type, token.token_value, sep='\t\t')

print('-'* 20)

new_lx = Lexer(source)
while new_lx.cur_ch != EOF:
    token = new_lx.next_token()
    print(token.token_type, token.token_value, sep='\t\t')


print('-'* 20)
for t in new_lx.token_list:
    print(t)
```