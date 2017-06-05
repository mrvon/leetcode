#include <stdio.h>

int majorityElement(int* nums, int numsSize) {
    int bound = numsSize / 2;
    int major = 0;
    int mask = 1;
    for (int i = 0; i < 32; i++) {
        int count = 0;
        for (int j = 0; j < numsSize; j++) {
            if (nums[j] & mask) {
                count++;
            }
            if (count > bound) {
                major |= mask;
                break;
            }
        }
        mask <<= 1;
    }
    return major;
}

int main() {
    int nums[] = {-2147483648};
    printf("%d\n", majorityElement(nums, sizeof(nums)/sizeof(int)));
}
