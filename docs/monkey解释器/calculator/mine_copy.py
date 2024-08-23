from collections import namedtuple
from typing import List

Token = namedtuple("Token", ["type", "val"])

white_space = set([" ", "\n", "\t"])

operators = {
    "+": "plus",
    "-": "minus",
    "*": "mul",
    "/": "div",
    "(": "lpar",
    ")": "rpar",
}


def tokenize(code: str) -> List[Token]:
    i, tokens = 0, []
    while i < len(code):
        if code[i] in white_space:
            i += 1
            continue
        if code[i] in operators:
            op = code[i]
            tokens.append(Token(op, op))
        if code[i].isnumeric():
            tmp = code[i]
            j = i + 1
            while j < len(code) and code[j].isnumeric():
                tmp += code[j]
                j += 1
            i = j - 1
            tokens.append(Token("num", int(tmp)))
        i += 1
    tokens.append(Token("eof", "eof"))
    return tokens


class ASTNode:
    pass


class Num(ASTNode):
    def __init__(self, token) -> None:
        self.token = token


class BinOp(ASTNode):
    def __init__(self, left, op, right):
        self.left = left
        self.op: Token = op
        self.right = right


class UnaryOp(ASTNode):
    """单目运算符"""

    def __init__(self, op, expr):
        self.op: Token = op
        self.expr = expr


"""
expr   : term ((PLUS | MINUS) term)*
term   : factor ((MUL | DIV) factor)*
factor : INTEGER | LPAREN expr RPAREN
"""


class Parser:
    def __init__(self, tokens: List[Token]) -> None:
        self.tokens = tokens
        if not tokens:
            self.cur = -1
        else:
            self.cur = 0

    @property
    def cur_token(self):
        return self.tokens[self.cur]

    def expr(self):
        """expr   : term ((PLUS | MINUS) term)*"""
        node = self.term()
        while self.cur_token.type in ("+", "-"):
            op = self.cur_token
            self.cur += 1
            node = BinOp(left=node, op=op, right=self.term())
        return node

    def factor(self):
        """factor : num | ( expr )"""
        if self.cur_token.type == "eof":
            return None
        if self.cur_token.type == "num":
            number = self.cur_token
            self.cur += 1
            return Num(number)
        if self.cur_token.type == "(":
            self.cur += 1
            node = self.expr()
            assert self.cur_token.type == ")"
            self.cur += 1
            return node
        if self.cur_token.type == '-':
            op = self.cur_token
            self.cur += 1
            node = UnaryOp(op=op, expr=self.factor())
            return node

    def term(self):
        """term   : factor ((MUL | DIV) factor)*"""
        node = self.factor()  # ?
        while self.cur_token.type in ("*", "/"):
            op = self.cur_token
            self.cur += 1
            node = BinOp(left=node, op=op, right=self.factor())
        return node


class Interpreter:
    def __init__(self, root) -> None:
        self.ast_root = root

    def visit(self, node):
        method_name = "visit_" + type(node).__name__
        visit_method = getattr(self, method_name)
        return visit_method(node)

    def visit_BinOp(self, node: BinOp):
        if node.op.val == "+":
            return self.visit(node.left) + self.visit(node.right)
        if node.op.val == "-":
            return self.visit(node.left) - self.visit(node.right)
        if node.op.val == "*":
            return self.visit(node.left) * self.visit(node.right)
        if node.op.val == "/":
            return self.visit(node.left) // self.visit(node.right)

    def visit_Num(self, node: Num):
        return int(node.token.val)
    
    def visit_UnaryOp(self, node: UnaryOp):
        op = node.op.type
        if op == '-':
            return -self.visit(node.expr)


def main():
    while True:
        text = input(">>>")
        if text in set(["quit", "exit", "q"]):
            return
        tokens = tokenize(text)

        parser = Parser(tokens)
        ast_tree = parser.expr()

        interp = Interpreter(ast_tree)
        res = interp.visit(ast_tree)
        print(res)


main()
