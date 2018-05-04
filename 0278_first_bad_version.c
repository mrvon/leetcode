#include <stdio.h>
#include <stdbool.h>

// Forward declaration of isBadVersion API.
static int global_version[] = {
    0,0,0,0,0,0,1,1,1,1,1,1,1,
};

bool isBadVersion(int version) {
    return global_version[version-1];
}

int firstBadVersion(int n) {
    int left = 1;
    int right = n;
    
    while (left < right) {
        int mid = left + (right - left) / 2;
        if (isBadVersion(mid)) {
            // mid may or may not the first bad version,
            // search for [left, mid]
            right = mid;
        } else {
            // [left, mid] is all good version.
            // search for [mid+1, right]
            left = mid+1;
        }
    }

    return left;
}

int main() {
    printf("%d\n", firstBadVersion(sizeof(global_version)/sizeof(int)));
}
