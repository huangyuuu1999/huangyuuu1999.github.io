#include<stdlib.h>
#include<stdio.h>

typedef struct IntSlice {
    int* Data;
    int Len, Cap;
} IntSlice;

IntSlice IntSliceInit(IntSlice s, int cap) {
    s.Cap = cap;
    s.Len = 0;
    s.Data = (int*)malloc(sizeof(int) * cap);
    return s;
}

int Len(IntSlice s) {
    return s.Len;
}

int Cap(IntSlice s) {
    return s.Cap;
}

void PrintSlice(IntSlice s) {
    int i;
    printf("[ ");
    for (i=0; i<s.Len; i++) {
        printf("%d ", s.Data[i]);
    } printf("] \n");
}

IntSlice append(IntSlice s, int elem) {
    if (s.Cap <= s.Len) {
        s.Cap = s.Cap * 2; // just for test
        int *new_data_address = (int*)malloc(sizeof(int) * s.Cap);
        for (int k=0; k<s.Len; k++) {
            new_data_address[k] = s.Data[k];
        }
        s.Data = new_data_address;    
    } // cap not enough
    s.Data[s.Len++] = elem;
    return s;
}

int DataAsAddress(IntSlice s) {
    return (int)(s.Data);
}

void test(IntSlice);

#define SLICE_INFO(s) \
PrintSlice(s);\
printf("len %d cap %d array address %x \n", Len(s), Cap(s), DataAsAddress(s));

int main() {
    IntSlice s;
    s = IntSliceInit(s, 2);
    // add by hand for easy
    s.Data[0] = 7;
    s.Data[1] = 6;
    s.Len = 2;
    SLICE_INFO(s) // [ 7 6 ]

    s = append(s, 3); // expand here
    SLICE_INFO(s)

    s = append(s, 42);
    SLICE_INFO(s) // [ 7 6 3 42 ] 
    
    // full test case1, expand in func
    test(s);
    SLICE_INFO(s)

    // cap 8, len 5, test case2, no expand in func
    s = append(s, 4399);
    SLICE_INFO(s)
    test(s);

    SLICE_INFO(s)
    return 0;
}

void test(IntSlice s) {
    printf("------ in  test ------\n");
    IntSlice s1 = append(s, 999); // s1 is a new struct obj
    SLICE_INFO(s1);
    printf("------ end test ------\n");
}