#include <stdio.h>
#include <stdlib.h>

void printArr(int* a, int n) {
    printf("[ ");
    for(int i=0; i<n; i++) {
        printf("%d ", a[i]);
    }
    printf("]\n");
}


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

void test_merge() {
    int a[] = {2, 4, 6, 8, 0, 1, 3, 9};
    int size = sizeof(a) / sizeof(int);

    merge(a, size, 4);
    printArr(a, size);
}

// 二分搜索: 在arr[lo..hi]内搜素元素target应该插入的位置
int bisect_left(int* arr, int lo, int hi, int target) {
    int i = lo, j = hi;
    while(i <= j) {
        int m = i + (j - i) / 2;
        if(arr[m] >= target) {
            j = m - 1;
        } else {
            i = m + 1;
        }
    }
    return i;
}

void test_bisect_left() {
    int a[] = {2, 4, 6, 8, 0, 1, 3, 9};
    int size = sizeof(a) / sizeof(int);
    int p = bisect_left(a, 4, size-1, 2);
    printf("bisect_left: %d\n", p);
}

// 对数组 arr[start..end]进行shift次 循环右移 操作
// 时间复杂度是 O(shift * (end - start))
void shiftRight(int* arr, int start, int end, int shift) {
    for(int s = 0; s<shift; s++) {
        int tmp = arr[end];
        for(int i=end; i>start; i--) {
            arr[i] = arr[i-1];
        }
        arr[start] = tmp;
    }
}

void testShiftRight() {
    printf("\n=== testShiftRight ===\n");
    int a[] = {1, 2, 3, 4, 5, 6};
    shiftRight(a, 0, 4, 2); // 451236
    printArr(a, 6);
}


// 数组a长度为n, a[0..k-1]是有序的，a[k..n-1]是有序的，merge使整个数组有序
void merge1(int* arr, int n, int k) {
    int i = 0, j = k; // 分别指向两个区域的第一个元素
    while(i < j && j < n) {
        int p = bisect_left(arr, j, n-1, arr[i]) - 1; // 寻找左侧的首位元素应该插入的位置
        shiftRight(arr, i, p, p-j+1);
        i += p - j + 2;
        j = p + 1;
    }
}

void test_merge1() {
    printf("\n=== test_merge1 ====\n");
    int a[] = {2, 4, 6, 8, 0, 1, 3, 9};
    int size = sizeof(a) / sizeof(int);

    merge1(a, size, 4);
    printArr(a, size);

    int b[] = {1, 2, 3, 4, 2, 5, 6};
    int sizeb = sizeof(b) / sizeof(int);
    merge1(b, sizeb, 4);
    printArr(b, sizeb);
}

void Swap(int* p, int* q) {
    int r = *p;
    *p = *q;
    *q = r;
}

// 简单插入排序，稳定，时间最坏O(n^2)
void insert_sort(int* arr, int n) {
    for(int i = 1; i<n; i++) { // 每一轮处理下标为i的数字
        for(int j=i; j>=1; j--) { // 依次检查是不是比前面的小，是的话就往前换
            if(arr[j] < arr[j-1]) {
                Swap(arr+j, arr + j-1);
            }
        }
        printArr(arr, n);
    }
}

// 简单插入排序，稳定，时间最坏O(n^2), 当已经有序的时候，时间O(n) 最好
void insert_sort2(int* arr, int n) {
    for(int i = 1; i<n; i++) { // 每一轮处理下标为i的数字
        for(int j=i; j>=1; j--) { // 依次检查是不是比前面的小，是的话就往前换
            if(arr[j] < arr[j-1]) {
                Swap(arr+j, arr + j-1);
            } else {
                break; // 遇到左边小于右边的时候就已经到了正确的位置，不需要再比较了
            }
        }
        printArr(arr, n);
    }
}

void test_insert_sort(void f(int*, int)) {
    printf("\n=== test_insert_sort ====\n");
    int a[] = {5, 1, 7, 3, 2, 4, 0, 6};
    int size = sizeof(a) / sizeof(int);
    printf("case: ");
    printArr(a, size);
    f(a, size);

    int b[] = {6, 2, 3, 5, 4, 1};
    int sizeb = sizeof(b) / sizeof(int);
    printf("case: ");
    printArr(a, size);
    f(b, sizeb);
}

int main() {
    test_merge();

    test_bisect_left();

    testShiftRight();

    test_merge1();

    test_insert_sort(insert_sort);

    test_insert_sort(insert_sort2);

    return 0;
}