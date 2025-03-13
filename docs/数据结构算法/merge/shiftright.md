# 右旋数组的多种方法

## 循环右移的方法
最坏是O(n^2)
```c
// 将数组的 num[s] ~ nums[t] 部分向右边移动k个位置
void shift_right(int* nums, int s, int t, int k) {
    for(int round=0; round < k; round++) {
        int tmp = nums[t];
        for(int j=t; j>s; j--) {
            nums[j] = nums[j-1];
        }
        nums[s] = tmp;
    }
}

void rotate(int* nums, int numsSize, int k) {
    k %= numsSize;
    shift_right(nums, 0, numsSize-1, k);
}
```

## 小优化，循环左移

```c
// 将数组的 num[s] ~ nums[t] 部分向 右边 移动k个位置
void shift_right(int* nums, int s, int t, int k) {
    for (int round = 0; round < k; round++) {
        int tmp = nums[t];
        for (int j = t; j > s; j--) {
            nums[j] = nums[j - 1];
        }
        nums[s] = tmp;
    }
}

// 将数组的 num[s] ~ nums[t] 部分向 左边 移动k个位置
void shift_left(int* nums, int s, int t, int k) {
    for (int round = 0; round < k; round++) {
        int tmp = nums[s];
        for (int i = s; i < t; i++) {
            nums[i] = nums[i + 1];
        }
        nums[t] = tmp;
    }
}

// 将数组的 nums[s] ~ nums[t] 位置的元素逆置
void reverse(int* a, int s, int t) {
    for (int i = s, j = t, tmp; i < j; i++, j--) {
        tmp = a[i];
        a[i] = a[j];
        a[j] = tmp;
    }
}

void rotate(int* nums, int numsSize, int k) {
    k %= numsSize;
    // 向右边移动k位置，等效于向左边移动n-k位置
    int right_shift = k;
    int left__shift = numsSize - k;
    if(right_shift < left__shift) {
        shift_right(nums, 0, numsSize-1, right_shift);
    } else {
        shift_left(nums, 0, numsSize-1, left__shift);
    }
}

```

## 三次逆置法
最坏也是小于 3O(n)
```c
// 将数组的 nums[s] ~ nums[t] 位置的元素逆置
void reverse(int* a, int s, int t) {
    for(int i = s, j = t, tmp; i < j; i++, j--) {
        tmp = a[i];
        a[i] = a[j];
        a[j] = tmp;
    }
}

void rotate(int* nums, int numsSize, int k) {
    k %= numsSize;
    // shift_right(nums, 0, numsSize-1, k);

    reverse(nums, numsSize-k, numsSize-1);
    reverse(nums, 0, numsSize-k-1);
    reverse(nums, 0, numsSize-1);
}

```

## 轮转数组
长度为n的数组，向右边移动k位，原先位于x下标的元素，现在到了(x+k) % n
会不会出现冲突的情况，例如多个值到了同样的新位置？不会，因为轮转前后的下标是一一对应的。

```c

//【轮转数组方法】新开辟一个数组，把旧的数字放到新的位置
void rotate_in_newarr(int* nums, int numsSize, int k) {
    int* tmp = (int*)malloc(sizeof(int) * numsSize);
    for(int i=0; i<numsSize; i++) {
        tmp[(i + k) % numsSize] = nums[i];
    }
    for(int i=0; i<numsSize; i++) {
        nums[i] = tmp[i];
    }
    free(tmp);
}

void rotate(int* nums, int numsSize, int k) {
    k %= numsSize;
    rotate_in_newarr(nums, numsSize, k);
}
```

## 轮转数组，不知道定理
不需要计算得到 “我们要循环多少次”，而是计数，等到全部的数字都被处理了，就退出。
```c
void without_gcd(int* nums, int numsSize, int k) {
    // assert k < numsSize
    int cnt = 0; // 处理过的数字
    int start = 0;
    while(cnt < numsSize) {
        int curIndex = start, curValue = nums[start];
        do {
            int next = (curIndex + k) % numsSize;
            int tmp = nums[next];
            nums[next] = curValue;
            curValue = tmp;
            curIndex = next;
            if(++cnt == numsSize) {
                return;
            }
        } while(curIndex != start);
        start++;
    }
}

void rotate(int* nums, int numsSize, int k) {
    k %= numsSize;
    without_gcd(nums, numsSize, k);
}
```
## 有gcd可以使用

使用数学直接计算出我们需要循环（外层）多少次。

```c
int gcd(int a, int b) {
    return b ? gcd(b, a % b) : a;
}

void with_gcd(int* nums, int numsSize, int k) {
    int limit = gcd(k, numsSize);
    for(int start=0; start<limit; start++) {
        int curIndex = start, curValue = nums[start];
        do {
            int next = (curIndex + k) % numsSize;
            int tmp = nums[next];
            nums[next] = curValue;
            curValue = tmp;
            curIndex = next;
        } while(start != curIndex);
    }
}

void rotate(int* nums, int numsSize, int k) {
    k %= numsSize;
    with_gcd(nums, numsSize, k);
}
```