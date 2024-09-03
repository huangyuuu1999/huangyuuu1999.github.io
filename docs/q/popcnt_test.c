#include <stdio.h>
#include <popcntintrin.h> // For _mm_popcnt_u32

int main() {
    unsigned int x = 5; // Binary: 101
    int count = _mm_popcnt_u32(x);
    printf("The number of 1's in %u is %d\n", x, count);
    return 0;
}