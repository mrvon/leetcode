/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
void deleteNode(struct ListNode* node) {
    if (node == NULL) {
        // THIS CASE IS IMPOSSIBLE
        return;
    } else if (node->next == NULL) {
        // THIS CASE IS IMPOSSIBLE
        return;
    } else {
        node->val = node->next->val;
        node->next = node->next->next;
    }
}
