// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

struct Solution;

impl Solution {
    pub fn merge_two_lists(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        match (list1, list2) {
            (None, None) => None,
            (Some(n), None) | (None, Some(n)) => Some(n),
            (Some(list1), Some(list2)) => {
                if list1.val >= list2.val {
                    Some(Box::new(ListNode{
                        val: list2.val,
                        next: Self::merge_two_lists(Some(list1), list2.next),
                    }))
                } else {
                    Some(Box::new(ListNode{
                        val: list1.val,
                        next: Self::merge_two_lists(list1.next, Some(list2)),
                    }))
                }
            }
        }
    }
}

fn main() {
    Solution::merge_two_lists(None, None);
}
