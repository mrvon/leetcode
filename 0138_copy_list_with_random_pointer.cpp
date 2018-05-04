/**
 * Definition for singly-linked list with a random pointer.
 * struct RandomListNode {
 *     int label;
 *     RandomListNode *next, *random;
 *     RandomListNode(int x) : label(x), next(NULL), random(NULL) {}
 * };
 */
class Solution {
    public:
        RandomListNode *copyRandomList(RandomListNode *head) {
            map<int, RandomListNode*> share;
            return copy(share, head);
        }

    private:
        RandomListNode *copy(map<int, RandomListNode*>& share, RandomListNode *head) {
            if (head == nullptr) {
                return nullptr;
            }
            auto newNode = share[head->label];
            if (newNode == nullptr) {
                newNode = new RandomListNode(head->label);
                share[head->label] = newNode;
                newNode->next = copy(share, head->next);
                newNode->random = copy(share, head->random);
            }
            return newNode;
        }
};
