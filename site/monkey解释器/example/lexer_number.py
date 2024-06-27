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

NUM = "NUM"  # 数字字面量

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
        word = self.source[word_start: self.cur_pos]  # 此时cur_pos对应的self.ch已经是非字母，是单词的下一位
        # 让当前读取位置 往前回一位
        self.cur_pos -= 1
        self.next_pos -= 1
        self.ch = self.source[self.cur_pos]
        return word
    
    def look_up_number(self):
        assert self.cur_ch.isnumeric()
        number_start = self.cur_pos
        while self.cur_ch.isnumeric():
            self.read_char()
        number = self.source[number_start: self.cur_pos]
        self.cur_pos -= 1
        self.next_pos -= 1
        self.ch = self.source[self.cur_pos]
        return number

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
        elif ch.isnumeric():
            number = self.look_up_number()
            token = Token(NUM, int(number))
            self.token_list.append(token)
        else:
            token = Token('Unknown')
        return token


source = "{(y}9let),;x+12 afuncbc=123 func("  # 注意空格也会被识别为unknown
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