from openai import OpenAI

api_key = ""

with open("apikey.txt") as f:
    api_key = f.read()

client = OpenAI(
    api_key=api_key,
    base_url="https://dashscope.aliyuncs.com/compatible-mode/v1",
)


completion = client.chat.completions.create(
    model="qwen2.5-7b-instruct-1m",
    messages=[
        {"role": "system", "content": "You are a helpful assistant."},
        {"role": "user", "content": "你认为叶文洁做的对吗？"},
    ],
)

print(completion.model_dump_json())
