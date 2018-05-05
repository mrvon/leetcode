// Brilliant!
// Show how to level-order traversal iteratively with constant extra space

/**
 * Definition for binary tree with next pointer.
 * struct TreeLinkNode {
 *  int val;
 *  TreeLinkNode *left, *right, *next;
 *  TreeLinkNode(int x) : val(x), left(NULL), right(NULL), next(NULL) {}
 * };
 */
class Solution {
    public:
        void connect(TreeLinkNode *root) {
            while (root != nullptr) {
                auto tempChild = new TreeLinkNode(0);
                auto currentChild = tempChild;
                while (root != nullptr) {
                    if (root->left != nullptr) {
                        currentChild->next = root->left;
                        currentChild = currentChild->next;
                    }
                    if (root->right != nullptr) {
                        currentChild->next = root->right;
                        currentChild = currentChild->next;
                    }
                    root = root->next;
                }
                root = tempChild->next;
            }
        }
};
