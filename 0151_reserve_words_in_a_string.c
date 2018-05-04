#include <stdio.h>
#include <string.h>

void removeSpace(char *s) {
    int i = 0;
    int j = strlen(s) - 1;
    int k = 0;
    int l = 0;

    while (i <= j) {
        if (s[i] == ' ') {
            i++;
        } else {
            break;
        }
    }

    while (i <= j) {
        if (s[j] == ' ') {
            j--;
        } else {
            break;
        }
    }

    if (i > j) {
        s[k] = 0;
    }

    int c = 0;
    for (k = i; k <= j; k++) {
        if (s[k] == ' ') {
            c++;
            if (c <= 1) {
                s[l++] = s[k];
            }
        } else {
            c = 0;
            s[l++] = s[k];
        }
    }
    s[l] = 0;
}

void reverseWord(char *s, int left, int right) {
    while (left < right) {
        char t = s[left];
        s[left] = s[right];
        s[right] = t;
        left++;
        right--;
    }
}

void reverseWords(char *s) {
    removeSpace(s);

    int left = 0;
    int right = strlen(s) - 1;

    while (left < right) {
        char t = s[left];
        s[left] = s[right];
        s[right] = t;
        left++;
        right--;
    }

    left = 0;
    right = strlen(s) - 1;

    int length = 0;

    while (left + length <= right) {
        if (s[left] == ' ') {
            left++;
        } else {
            length++;
            if (s[left+length] == ' ') {
                reverseWord(s, left, left+length-1);
                left = left+length;
                length = 0;
            }
        }
    }

    if (length > 0) {
        reverseWord(s, left, left+length-1);
    }
}

int main() {
    char s[] = " the sky is blue ";
    reverseWords(s);
    printf("%s\n", s);
}
