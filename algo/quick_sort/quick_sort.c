// Copy from K&R book
void
swap(int v[], int i, int j) {
    int temp;
    temp = v[i];
    v[i] = v[j];
    v[j] = temp;
}

void
qsort(int v[], int left, int right) {
    int i;
    int last;
    if (left >= right) {
        return;
    }
    swap(v, left, (left + right) / 2);
    last = left;
    for (i = left + 1; i <= right; i++) {
        if (v[i] < v[left]) {
            swap(v, ++last, i);
        }
    }
    swap(v, left, last);
    qsort(v, left, last - 1);
    qsort(v, last + 1, right);
}

#include <stdio.h>
int
main() {
    int t[] = {9, 11, 0, 3, 1, 3, 4, 5, 6};
    int len = sizeof(t) / sizeof(int);
    qsort(t, 0, len - 1);
    for (int i = 0; i < len; i++) {
        printf("%d ", t[i]);
    }
    printf("\n");
}
