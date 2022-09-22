struct Solution;

impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        if x == 0 {
            return true;
        } else if x < 0 {
            return false;
        }
        let s = x.to_string();
        let bytes = s.as_bytes();
        for (i, c) in bytes.iter().enumerate() {
            let j = bytes.len() - 1 - i;
            if i >= j {
                return true;
            }
            if c != &bytes[j] {
                return false;
            }
        }
        return true;
    }
}

fn main() {
    assert!(Solution::is_palindrome(0) == true);
    assert!(Solution::is_palindrome(-1) == false);
    assert!(Solution::is_palindrome(1) == true);
    assert!(Solution::is_palindrome(11) == true);
    assert!(Solution::is_palindrome(121) == true);
    assert!(Solution::is_palindrome(123) == false);
    assert!(Solution::is_palindrome(-2147483648) == false);
}
