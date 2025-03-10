# 1. 定义一个 UTF-8 的替换字符（�）的字节流
original_bytes = b'\xEF\xBF\xBD'  # UTF-8 的替换字符 �

# 2. 错误地用 GBK 编码解码这段字节
decoded_text = original_bytes.decode('GBK', errors='ignore')  # 强制用 GBK 解码

# 3. 输出结果
print(decoded_text)  # 输出：锟