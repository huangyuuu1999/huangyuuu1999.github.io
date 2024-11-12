from collections import namedtuple
from typing import List

Token = namedtuple("Token", ["type", "val"])

white_space = set([" ", "\n", "\t"])

operators = {
    # 运算符 一词一码
    "+": "plus",
    "-": "minus",
    "*": "mul",
    "/": "div",
}


class Error(Exception):
    def __init__(self, error_code=None, message=None):
        self.error_code = error_code
        self.message = f"{self.__class__.__name__}: {message}"


class ParserError(Error):
    def __init__(self, error_code=None, message=None):
        super().__init__(error_code, message)


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
    """双目运算符"""

    def __init__(self, left, op, right):
        self.left = left
        self.op: Token = op
        self.right = right


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
        """factor : INTEGER | LPAREN expr RPAREN"""
        if self.cur_token.type == "eof":
            return None
        if self.cur_token.type == "num":
            number = self.cur_token
            self.cur += 1
            return Num(number)

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
