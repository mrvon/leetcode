/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode *getIntersectionNode(struct ListNode *headA, struct ListNode *headB) {
    int lenA = 0;
    int lenB = 0;

    for (struct ListNode *p = headA; p != NULL; p = p->next) {
        lenA++;
    }

    for (struct ListNode *p = headB; p != NULL; p = p->next) {
        lenB++;
    }

    if (lenA > lenB) {
        int tl = lenA;
        lenA = lenB;
        lenB = tl;
        struct ListNode* tp = headA;
        headA = headB;
        headB = tp;
    }

    int diff = lenB - lenA;
    for (int i = 0; i < diff; i++) {
        headB = headB->next;
    }

    while (headB) {
        if (headB == headA) {
            return headB;
        }
        headB = headB->next;
        headA = headA->next;
    }

    return NULL;
}
