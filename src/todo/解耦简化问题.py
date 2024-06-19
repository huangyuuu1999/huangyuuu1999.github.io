class B:
    get_1 = 1
    get_2 = 2

    def func1(self):
        return 1

    def func2(self):
        return 2


for item in dir(B):
    if item.startswith('get'):
        original = getattr(B, item)
        setattr(B, item, original ** 2)

print(B.get_2, B.get_1)


for func in dir(B):
    if func.startswith('func') and callable(getattr(B, func)):
        original_func = getattr(B, func)
        print(original_func(B))
        
        def new_func(self):
            print('新的函数')
            ret = original_func(self)
            return ret
        setattr(B, func, new_func)
b = B()
print(b.func2())
print(b.func1())
