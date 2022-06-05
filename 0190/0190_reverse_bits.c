#include <stdio.h>
#include <stdint.h>

uint32_t reverseBits(uint32_t n) {
    uint32_t m = 0;
    for (int i = 0; i < 32; i++) {
        if (n & 0x1) {
            m = m | 0x1;
        }
        n = n >> 1;
        if (i < 31) {
            m = m << 1;
        }
    }
    return m;
}

int main() {
    printf("%u\n", reverseBits(43261596));
}
