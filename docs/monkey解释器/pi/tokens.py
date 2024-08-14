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
    "(": "lpar",
    ")": "rpar",
}


def parse(code: str) -> List[Token]:
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
        # else:
        #     tokens.append(Token('?', '#'))
        i += 1
    return tokens
