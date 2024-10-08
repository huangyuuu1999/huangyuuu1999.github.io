# 词法解析

无论是编译型还是解释型语言，为了能够把编程语言转化为计算机能理解的内容，首先就要对用户输入的内容进行拆分，拆分的结果就是token。
完成这件 **识别、拆分** 任务的代码模块，就是词法解析器，是一种处理文本（符合特定语法规定的）的工具。

## 词法解析器
词法分析器往往是编译或者解释的第一步，他将源代码抽象成为token序列。
词法分析器（lexer）也叫词法单元生成器（tokenizer）或者扫描器（scanner）。

### 词法分析器的功能
词法分析器是一些代码，它们能够实现这样的功能：把用户写的源代码，转换成一堆词法单元（token）。就像是你在翻译文言文的时候，首先需要把一句话拆分成为许多不可再分的单元。

举个例子：你要翻译“庭有枇杷树”，首先就要把这句话拆分为 `["庭","有","枇杷","树"]` 或者 `["庭","有","枇杷树"]`，你要做的事情就是拆分，得到一个单词（有的单词就是一个字）序列，而**不必关心**这句话表达什么意思，是不是符合古汉语的语法（这是后面的语法分析和语义分析要做的事情）。

### 词法解析小案例
对于一门像样的编程语言写的代码，例如下面的这一句：
```js
// 注意这不是任何一种已经存在的语言，
// 只是我自己规定的假想象一种语言的语法，尽管他看上去有许多其他语言的影子
let x = 5 + 5;
```
这门语言的词法解析器应该能辨认出：
最开始有一个关键字let，然后是一个变量的名字x，后面是赋值符号=，加法符号+，还有两个数字5。
词法解析器应该按顺序给我们这样的结果（token序列）

```python
[
    LET, # 关键字，语言规定的，例如if else while struct await func 等
    IDENTIFIER('x'), # 标识符 变量名x, abc student_age, 函数名 add, push_back
    EQUAL_SIGN, # 符号 + - * / % & > < <= || ? 等
    INTEGER(5), # 数字字面量 
    PLUS_SIGN, 
    INTEGER(5),
]
```

这个结果中的每一个成员，都是一个词法单元token，每一个词法单元，都在源代码中有对应。每一个词法单元token都有自己的类型，主要有这几类

- 关键字 ——— 由一门语言自己规定，不同的语言关键字表不同，例如`c语言`的`include typedef if else while goto` 
- 运算符 加`+`减`-`乘`*`除`/`，取模`%`，与`&`或`|`非`^`关系运算，大于等于`>=`, 有的语言还有三元运算`？`符号，指针取值符`->`，解引用符`*`等
- 界限符 `{} () [] <>` 等, 用于在语句中划分出小的作用域，或者数组根据下标取值等
- 标识符 变量名`x y z abc p_name`，函数名`func add getEventLoop`等
- 字面量 写一个数字 `42, 0, 0xfff, 3.24` 或者一个字符串 `"man out!"`

这些类型的token再加上一些空格换行，注释等，大概就是一门语言中你能写的所有东西。

## 总结
词法解析器的功能和目标：将符合某一门语言规则的文本源代码，转换成token序列。

token序列的种类：关键字、标识符、界限符、运算符、字面量。