# merge操作

## 使用额外的空间 O(n)
自己分配额外的空间，两个子数组的长度之和长度，然后双指针。
```c
// 数组a长度为n, a[0..k-1]是有序的，a[k..n-1]是有序的，merge使整个数组有序
void merge(int *arr, int n, int k) {
    int* tmp = (int*)malloc(sizeof(int) * n);
    int i = 0, j = k, p = 0;
    while (i < k && j < n) {
        if(arr[i] < arr[j]) {
            tmp[p++] = arr[i++];
        } else {
            tmp[p++] = arr[j++];
        }
    }
    while(i < k) {
        tmp[p++] = arr[i++];
    }
    while(j < n) {
        tmp[p++] = arr[j++];
    }
    for(int j=0; j<n; j++) {
        arr[j] = tmp[j];
    }
}

void printArr(int* a, int n) {
    printf("[ ");
    for(int i=0; i<n; i++) {
        printf("%d ", a[i]);
    }
    printf("]\n");
}

void test_merge() {
    int a[] = {2, 4, 6, 8, 0, 1, 3, 9};
    int size = sizeof(a) / sizeof(int);

    merge(a, size, 4);
    printArr(a, size);
}
```

https://leetcode.cn/problems/sorted-merge-lcci/description/
lc上 合并数组那道题 跟这个也差不多，只不过是从后往前填写。

## 不使用额外的空间，直接排序
直接使用任意一种原地排序算法：冒泡，选择，插入，快速排序等等。

插入排序对于已经有序的，有一点优势。最好O(n)，最坏O(n^2)。

对于后半段的，依次往前插，n-k个，最坏 k (n - k)。

### 不使用额外空间，快排 
选择最好O(n) 最坏O(n^2), 快速最好O(nlogn)。

## 使用教材上的 循环右移方法
最坏也是O(n^2)

## 还有O(nlogn)的方法
