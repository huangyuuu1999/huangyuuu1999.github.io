给C添加属性，表现为深拷贝，添加函数成员，是否也是深拷贝？
```python
class C:
    pass

lst = [1, 2, 4]

setattr(C, 'lst', lst)

print(C.lst)
lst = [1, 5, 7]

print(C.lst)


def func():
    return 1

setattr(C, 'f', func)

print(C.f())

def func():
    return 5
setattr(C, 'f', func)

print(C.f())

# [1, 2, 4]
# [1, 2, 4]
# 1
# 5
```