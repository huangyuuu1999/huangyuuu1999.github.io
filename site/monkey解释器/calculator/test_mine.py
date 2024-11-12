from mine import parse, Token


def test_parse():
    code = "( 17 +26 * 4) - 06"  # 06 算 6
    tokens = parse(code)
    for token in tokens:
        print(token.type, "\t", token.val)
    assert len(tokens) == 9


def test_parse_valid():
    code = "12 + 7( ) 8U -l 5 /1 -0"
    tokens = parse(code)
    for token in tokens:
        print(token.type, "\t", token.val)
    assert Token("num", 0) in tokens
    assert Token("num", 8) in tokens


def test_empty():
    code = ""
    tokens = parse(code)
    assert tokens == []


def test_no_valid():
    code = "adnalbfwpqoubpf+++++***&^^"
    tokens = parse(code)
    assert len(tokens) == 8