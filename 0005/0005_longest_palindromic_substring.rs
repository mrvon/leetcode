struct Solution;

impl Solution {
    pub fn longest_palindrome(s: String) -> String {
        let s = s.as_bytes();
        let mut longest = "".as_bytes();
        for i in 0..s.len() {
            let r = Self::expand(s, i as i32);
            if r.len() > longest.len() {
                longest = r;
            }
        }
        String::from_utf8_lossy(longest).to_string()
    }

    fn expand(s: &[u8], p: i32) -> &[u8] {
        let mut s1: &[u8] = "".as_bytes();
        // p && p+1 as pair
        if (p as usize) < s.len() - 1 && s[p as usize] == s[(p + 1) as usize] {
            let mut l = p;
            let mut r = p + 1;
            while l >= 0 && (r as usize) < s.len() && s[l as usize] == s[r as usize] {
                l -= 1;
                r += 1;
            }
            s1 = &s[(l + 1) as usize..r as usize];
        }
        // p as middle single one
        let mut l = p - 1;
        let mut r = p + 1;
        while l >= 0 && (r as usize) < s.len() && s[l as usize] == s[r as usize] {
            l -= 1;
            r += 1;
        }
        let s2 = &s[(l + 1) as usize..r as usize];
        if s1.len() >= s2.len() {
            s1
        } else {
            s2
        }
    }
}

fn main() {
    assert_eq!("", Solution::longest_palindrome("".to_string()));
    assert_eq!("a", Solution::longest_palindrome("a".to_string()));
    assert_eq!("a", Solution::longest_palindrome("ab".to_string()));
    assert_eq!("c", Solution::longest_palindrome("cab".to_string()));
    assert_eq!("bb", Solution::longest_palindrome("abbd".to_string()));
    assert_eq!("cbbc", Solution::longest_palindrome("acbbcd".to_string()));
}
