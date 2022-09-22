// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode {
            next: None,
            val
        }
    }
}

struct Solution;

impl Solution {
    pub fn add_two_numbers(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut dummy_head = ListNode::new(0);
        let mut current = &mut dummy_head;
        let mut p = list1;
        let mut q = list2;
        let mut carry: i32 = 0;

        while p != None || q != None {
            let sum = match (&p, &q) {
                (Some(list1), Some(list2)) => list1.val + list2.val + carry,
                (Some(list1), None) => list1.val + carry,
                (None, Some(list2)) => list2.val + carry,
                (None, None) => carry,
            };
            carry = sum / 10;
            current.next = Some(Box::new(ListNode::new(sum % 10)));
            current = current.next.as_mut().unwrap();
            p = if p != None { p.unwrap().next } else { p };
            q = if q != None { q.unwrap().next } else { q };
        }
        if carry > 0 {
            current.next = Some(Box::new(ListNode::new(carry)))
        }
        dummy_head.next
    }
}

fn main() {
    Solution::add_two_numbers(None, None);
}
