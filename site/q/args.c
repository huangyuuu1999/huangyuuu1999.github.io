#include <stdarg.h>
#include <stdio.h>


double average(int param_num, ...) {
    va_list variable_args_list;
    double sum = 0.0;
    int index;
    /* 为 num 个参数初始化 valist */
    va_start(variable_args_list, param_num);

    /* 每次读一个参数 */
    for (index=0; index<param_num; index++) {
        sum += va_arg(variable_args_list, int);
    }
    /* 清理为 valist 保留的内存 */
    va_end(variable_args_list);
    return sum/param_num;
}

int main() {
    printf("Average of 2, 3, 4, 5 = %f\n", average(4, 2,3,4,5));
    printf("Average of 5, 10, 15 = %f\n", average(3, 5,10,15));
    return 0;
}