from functools import wraps


def print_before_get(cls):
    print(dir(cls))
    for func in dir(cls):
        if func.startswith('get') and callable(getattr(cls, func)):
            origin_func = getattr(cls, func)

            @wraps(origin_func)
            def wrapper(*args, **kwargs):
                print("类的装饰器")
                result = origin_func(*args, **kwargs)
                return result

            setattr(cls, func, wrapper)
        elif func.startswith('get'):
            setattr(cls, func, '666')
    print(dir(cls))
    return cls


def log_func(func):
    @wraps(func)
    def wrapped(*args, **kwargs):
        print("函数的装饰器")
        ret = func(*args, **kwargs)
        return ret

    return wrapped


@print_before_get
class A:
    get_3 = 3
    get_4 = 4

    # @staticmethod
    def get_1(self):
        return 1

    # @log_func
    # @staticmethod
    def get_2(self):
        return 2


if __name__ == '__main__':
    a = A()
    print(dir(a))
    # print(A.get_1())
    # print(A.get_1())
    print(a.get_1())
    print(a.get_2())
    print(A.get_3)
    print(A.get_4)
