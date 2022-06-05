/*
Since struct TreeNode hasn't parent pointer, we use stack to solve this problem.
*/

/**
 * Definition for binary tree
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */

typedef struct {
    struct TreeNode** buf;
    int top;
} Stack;

void stackCreate(Stack *stack, int maxSize) {
    stack->buf = malloc(sizeof(struct TreeNode*) * maxSize);
    stack->top = 0;
}

void stackPush(Stack *stack, struct TreeNode* element) {
    stack->buf[stack->top] = element;
    stack->top++;
}

void stackPop(Stack *stack) {
    stack->top--;
}

struct TreeNode* stackPeek(Stack *stack) {
    return stack->buf[stack->top-1];
}

bool stackEmpty(Stack *stack) {
    if (stack->top == 0) {
        return true;
    } else {
        return false;
    }
}

void stackDestroy(Stack *stack) {
    free(stack->buf);
}

struct BSTIterator {
    Stack s;
};

struct BSTIterator *bstIteratorCreate(struct TreeNode *root) {
    struct BSTIterator *iter = malloc(sizeof(struct BSTIterator));
    stackCreate(&iter->s, 1024);

    struct TreeNode* left = root;
    while (left) {
        stackPush(&iter->s, left);
        left = left->left;
    }

    return iter;
}

/** @return whether we have a next smallest number */
bool bstIteratorHasNext(struct BSTIterator *iter) {
    return !stackEmpty(&iter->s);
}

/** @return the next smallest number */
int bstIteratorNext(struct BSTIterator *iter) {
    struct TreeNode* node = stackPeek(&iter->s);
    stackPop(&iter->s);

    struct TreeNode* left = node->right;
    while (left) {
        stackPush(&iter->s, left);
        left = left->left;
    }

    return node->val;
}

/** Deallocates memory previously allocated for the iterator */
void bstIteratorFree(struct BSTIterator *iter) {
    stackDestroy(&iter->s);
    free(iter);
}

/**
 * Your BSTIterator will be called like this:
 * struct BSTIterator *i = bstIteratorCreate(root);
 * while (bstIteratorHasNext(i)) printf("%d\n", bstIteratorNext(i));
 * bstIteratorFree(i);
 */
