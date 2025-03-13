# 添加导入random模块
import random
from datetime import datetime

from openai import OpenAI

api_key = ""

with open("apikey.txt") as f:
    api_key = f.read()

## 步骤1:定义工具函数


# 模拟天气查询工具。返回结果示例：“北京今天是雨天。”
def get_current_weather(arguments):
    # 定义备选的天气条件列表
    weather_conditions = ["晴天", "多云", "雨天"]
    # 随机选择一个天气条件
    random_weather = random.choice(weather_conditions)
    # 从 JSON 中提取位置信息
    location = arguments["location"]
    # 返回格式化的天气信息
    return f"{location}今天是{random_weather}。"


# 查询当前时间的工具。返回结果示例：“当前时间：2024-04-15 17:15:18。“
def get_current_time():
    # 获取当前日期和时间
    current_datetime = datetime.now()
    # 格式化当前日期和时间
    formatted_time = current_datetime.strftime("%Y-%m-%d %H:%M:%S")
    # 返回格式化后的当前时间
    return f"当前时间：{formatted_time}。"


def save_info(content):
    with open("db.txt", "w+") as f:
        f.write(content)
    print("write ok")


# 请将以下代码粘贴到步骤1代码后

## 步骤2:创建 tools 数组

tools = [
    {
        "type": "function",
        "function": {
            "name": "get_current_time",
            "description": "当你想知道现在的时间时非常有用。",
        },
    },
    {
        "type": "function",
        "function": {
            "name": "get_current_weather",
            "description": "当你想查询指定城市的天气时非常有用。",
            "parameters": {
                "type": "object",
                "properties": {
                    "location": {
                        "type": "string",
                        "description": "城市或县区，比如北京市、杭州市、余杭区等。",
                    }
                },
                "required": ["location"],
            },
        },
    },
    {
        "type": "function",
        "function": {
            "name": "save_info",
            "description": "保存单词",
            "parameters": {
                "type": "object",
                "properties": {
                    "location": {
                        "type": "string",
                        "description": "城市或县区，比如北京市、杭州市、余杭区等。",
                    }
                },
                "required": ["location"],
            },
        },
    },
]
tool_name = [tool["function"]["name"] for tool in tools]
print(f"创建了{len(tools)}个工具，为：{tool_name}\n")


# 初始化一个 messages 数组
messages = [
    {
        "role": "system",
        "content": """你是一个很有帮助的助手。
        如果用户提问关于天气的问题，请调用 ‘get_current_weather’ 函数;
        如果用户提问关于时间的问题，请调用‘get_current_time’函数。
     请以友好的语气回答问题。""",
    },
    {"role": "user", "content": "上海天气"},
]

client = OpenAI(
    api_key=api_key,
    base_url="https://dashscope.aliyuncs.com/compatible-mode/v1",
)


def function_calling():
    # 模型列表：https://help.aliyun.com/zh/model-studio/getting-started/models
    completion = client.chat.completions.create(
        model="qwen2.5-7b-instruct-1m", messages=messages, tools=tools
    )
    print("返回对象：")
    print(completion.choices[0].message.model_dump_json())
    print("\n")
    return completion


print("正在发起function calling...")
completion = function_calling()
# 步骤5:运行工具函数
# 请将以下代码粘贴到步骤4 代码后
import json

print("正在执行工具函数...")
# 从返回的结果中获取函数名称和入参
function_name = completion.choices[0].message.tool_calls[0].function.name
arguments_string = completion.choices[0].message.tool_calls[0].function.arguments

# 使用json模块解析参数字符串
arguments = json.loads(arguments_string)
# 创建一个函数映射表
function_mapper = {
    "get_current_weather": get_current_weather,
    "get_current_time": get_current_time,
}
# 获取函数实体
function = function_mapper[function_name]
# 如果入参为空，则直接调用函数
if arguments == {}:
    function_output = function()
# 否则，传入参数后调用函数
else:
    function_output = function(arguments)
# 打印工具的输出
print(f"工具函数输出：{function_output}\n")
