from openai import OpenAI

api_key = ""

with open("apikey.txt") as f:
    api_key = f.read()


def get_response(messages):
    client = OpenAI(
        api_key=api_key,
        base_url="https://dashscope.aliyuncs.com/compatible-mode/v1",
    )
    # 模型列表：https://help.aliyun.com/zh/model-studio/getting-started/models
    completion = client.chat.completions.create(
        model="qwen2.5-7b-instruct-1m", messages=messages
    )
    return completion


# 初始化一个 messages 数组
messages = [
    {
        "role": "system",
        "content": """
        你是一个英语词典专家，对于用户给出的单词，你会回复其中文意思
        对于用户输入的 word 你的输出是output, 你使用markdown格式回复:
        - {{word}}
        - {{output}}
        例如 用户输入 cat 你的返回是:
        - cat
        - 猫
        """,
    }
]
assistant_output = "你可以输入任何英文单词，我会为你给出他的中文解释"
print(f"模型输出：{assistant_output}\n")

while True:
    user_input = input("请输入：")
    if user_input == "退出":
        break
    # 将用户问题信息添加到messages列表中
    messages.append({"role": "user", "content": user_input})
    assistant_output = get_response(messages).choices[0].message.content
    # 将大模型的回复信息添加到messages列表中
    messages.append({"role": "assistant", "content": assistant_output})
    print(f"模型输出:\n{assistant_output}")
    print("\n")
