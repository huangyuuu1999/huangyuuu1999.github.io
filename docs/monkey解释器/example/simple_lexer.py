source = """let x =2+ 7 ;
x =x -3;
"""

class Lexer:
    def __init__(self, source_code):
        self.source = source_code
        self.cur_pos = -1
        self.next_pos = 0
        self.cur_ch = ''

    def read_char(self):
        if self.next_pos >= len(self.source):
            self.cur_ch = 'EOF'
        else:
            self.cur_ch = self.source[self.next_pos]
        self.cur_pos = self.next_pos
        self.next_pos += 1
        return self.cur_ch

lx = Lexer(source)
print('ch', 'cur_pos', 'next_pos', sep='\t')
print('start', lx.cur_pos, lx.next_pos, sep='\t')

for _ in range(len(source)):
    c = lx.read_char()
    if c == '\n':
        print(r'\n', lx.cur_pos, lx.next_pos, sep='\t')
    else:    
        print(c, lx.cur_pos, lx.next_pos, sep='\t')

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
print(read_ch)


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