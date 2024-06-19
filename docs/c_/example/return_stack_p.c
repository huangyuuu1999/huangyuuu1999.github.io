// return_stack_p.c
#include<stdio.h>

int* f(int tmp) {
    int a = tmp;
    return &a;
}

int main() {
    int *b = f(1);
    printf("b is %d\n", *b); // b is 1
    f(2);
    printf("b is %d\n", *b); // b is 2
}

//return_stack_p.c:6:13: warning: address of stack memory associated with local variable 'a' returned [-Wreturn-stack-address]