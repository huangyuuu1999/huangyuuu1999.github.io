import asyncio

async def test_go_routine():
    print("hello goroutine!")
    return 42


async def main():
    res = await test_go_routine()
    print("in main: ", res)

if __name__ == '__main__':
    asyncio.run(main())

# hello goroutine!
# in main:  42