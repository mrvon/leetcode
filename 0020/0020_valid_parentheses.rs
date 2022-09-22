struct Solution;

impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut stack = Vec::new();
        for c in s.chars() {
            stack.push(match c {
                '(' => ')',
                '[' => ']',
                '{' => '}',
                _ => {
                    if Some(c) == stack.pop() {
                        continue;
                    } else {
                        return false;
                    }
                }
            })
        }
        stack.is_empty()
    }
}

fn main() {
    assert!(Solution::is_valid("{}".to_string()));
    assert!(Solution::is_valid("{}[]".to_string()));
    assert!(Solution::is_valid("(){}[]".to_string()));
    assert!(!Solution::is_valid("{]".to_string()));
    assert!(!Solution::is_valid("({])".to_string()));
    assert!(Solution::is_valid("({()})".to_string()));
}
