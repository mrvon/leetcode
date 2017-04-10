/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */

int depth(struct TreeNode* root) {
    if (root == NULL) {
        return 0;
    }

    return depth(root->left) + 1;
}

int countNodes(struct TreeNode* root) {
    if (root == NULL) {
        return 0;
    }

    int left_sub = depth(root->left);
    int right_sub = depth(root->right);

    if (left_sub == right_sub) {
        // left subtree is a full binary tree, right subtree is a complete
        // binary tree.
        // it's node number is that: 2^0 + 2^1 + 2^2 + ... + 2^depth-1
        // = 2^depth - 1
        int left_nodes = (1 << left_sub) - 1;
        return 1 + left_nodes + countNodes(root->right);
    } else {
        // right subtree is a full binary tree, left subtree is a complete
        // binary tree.
        int right_nodes = (1 << right_sub) - 1;
        return 1 + countNodes(root->left) + right_nodes;
    }
}
