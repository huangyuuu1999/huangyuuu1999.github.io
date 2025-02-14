# test_ord.py
# 使用 python 的 ord 函数 来获取字符在 unicode 字符集中的编号

emoji = "😅"

chinese_char = "你"

emoji_with_2_code_point = "💪🏿"

unicode_id1 = ord(emoji)
unicode_id2 = ord(chinese_char)
unicode_id3 = ord(emoji_with_2_code_point)

print(unicode_id1, unicode_id2, unicode_id3)

# ❯ python test_ord.py
# 128517 20320