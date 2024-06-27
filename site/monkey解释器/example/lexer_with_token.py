class Token:
    def __init__(self, token_type, token_value=None) -> None:
        self.token_type = token_type
        self.token_value = token_value

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

IDENT = 'IDENT' # 标识符 abc, add, foobar, x, y_1


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
    
    def next_token(self) -> Token:  # 目前只能识别运算符号+=和 界限符号,;(){}
        ch = self.read_char()
        if ch == 'EOF':
            token = Token(EOF)
        elif ch == '=':
            token = Token(ASSIGN)
        elif ch == '+':
            token = Token(PLUS)
        elif ch == ',':
            token = Token(COMMA)
        elif ch == ';':
            token = Token(SEMICOLON)
        elif ch == '(':
            token = Token(LPARENT)
        elif ch == ')':
            token = Token(RPARENT)
        elif ch == '{':
            token = Token(LBRACE)
        elif ch == '}':
            token = Token(RBRACE)
        else:
            token = Token('Unknown', ch)
        return token


source = "{(y}),;x+abc= ("  # 注意空格也会被识别为unknown
lx = Lexer(source)
for _ in range(len(source)+2):
    token = lx.next_token()
    print(token.token_type, token.token_value, sep='\t')