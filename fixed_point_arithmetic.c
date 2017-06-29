#include <stdio.h>
#include <assert.h>

// SCALE 16
// 1010 1010 1010 1010 . 1010 1010 1010 1010
// 
// SCALE 8
// 1010 1010 1010 1010 1010 1010 . 1010 1010
// 
// SCALE 24
// 1010 1010 . 1010 1010 1010 1010 1010 1010

#define SCALE           16
#define FRACTION_MASK   (0xffffffff >> (32 - SCALE))
#define WHOLE_MASK      (0xffffffff ^ FRACTION_MASK)

int double_to_fixed(double x) {
    return (x * (double)(1 << SCALE));
}

double fixed_to_double(int x) {
    return ((double)x / (double)(1 << SCALE));
}

int int_to_fixed(int x) {
    return x << SCALE;
}

int fixed_to_int(int x) {
    return x >> SCALE;
}

int fraction_part(int x) {
    assert(x >= 0);
    return x & FRACTION_MASK;
}

int whole_part(int x) {
    assert(x >= 0);
    return x & WHOLE_MASK;
}


int main() {
    int f1 = double_to_fixed(5.6);
    int f2 = double_to_fixed(7.9);
    int f3 = double_to_fixed(27.6789);

    printf("%lf + %lf = %lf\n", fixed_to_double(f1), fixed_to_double(f2), fixed_to_double(f1 + f2));
    printf("%lf - %lf = %lf\n", fixed_to_double(f1), fixed_to_double(f2), fixed_to_double(f1 - f2));
    printf("Fractional Part: %lf\n", fixed_to_double(fraction_part(f1)));
    printf("Fractional Part: %lf\n", fixed_to_double(fraction_part(f3)));
    printf("Whole Part: %lf\n", fixed_to_double(whole_part(f3)));
    printf("Min Number: %lf\n", fixed_to_double(1));
}
