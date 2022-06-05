// In-place solution
#include <stdio.h>
#include <stdlib.h>

#define IsLive(s) (s&0x1)

int liveNeighbours(int** board, int boardRowSize, int boardColSize, int x, int y) {
    int offset[8][2] = {
        {-1, -1},
        {-1, 0},
        {-1, 1},
        {0, -1},
        {0, 1},
        {1, -1},
        {1, 0},
        {1, 1},
    };

    int count = 0;
    for (int k = 0; k < 8; k++) {
        int i = x + offset[k][0];
        int j = y + offset[k][1];
        if (i >= 0 && i < boardRowSize && j >= 0 && j < boardColSize) {
            if (IsLive(board[i][j])) {
                count++;
            }
        }
    }
    return count;
}

void gameOfLife(int** board, int boardRowSize, int boardColSize) {
    for (int i = 0; i < boardRowSize; i++) {
        for (int j = 0; j < boardColSize; j++) {
            int n = liveNeighbours(board, boardRowSize, boardColSize, i, j);
            if (IsLive(board[i][j])) {
                if (n < 2) { // dies as if by under-population
                    // do nothing
                } else if (n > 3) { // dies as if by over-population
                    // do nothing
                } else { // keep it live
                    board[i][j] = 0x3; // 11b
                }
            } else {
                if (n == 3) { // reproduction
                    board[i][j] = 0x2; // 10b
                }
            }
        }
    }

    for (int i = 0; i < boardRowSize; i++) {
        for (int j = 0; j < boardColSize; j++) {
            board[i][j] >>= 1;
        }
    }
}

int print(int** board, int boardRowSize, int boardColSize) {
    for (int i = 0; i < boardRowSize; i++) {
        for (int j = 0; j < boardColSize; j++) {
            printf("%d ", board[i][j]);
        }
        printf("\n");
    }
    printf("\n");
}

int** build(int boardRowSize, int boardColSize) {
    int** board = malloc(sizeof(int*) * boardRowSize);
    for (int i = 0; i < boardRowSize; i++) {
        board[i] = malloc(sizeof(int) * boardColSize);
        for (int j = 0; j < boardColSize; j++) {
            if (rand() % 2) {
                board[i][j] = 0;
            } else {
                board[i][j] = 1;
            }
        }
    }
    return board;
}

int main() {
    int** board = build(8, 8);
    print(board, 8, 8);

    gameOfLife(board, 8, 8);
    print(board, 8, 8);

    gameOfLife(board, 8, 8);
    print(board, 8, 8);

    return 0;
}
