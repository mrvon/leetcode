struct Solution;

impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        use std::collections::HashMap;
        let s: Vec<char> = s.chars().collect();
        let mut unique: HashMap<char, bool> = HashMap::new();
        let mut max = 0;
        let mut len = 0;
        let mut right = 0;
        let mut left = 0;
        while right < s.len() {
            let c = s.get(right).unwrap();
            match unique.get(c) {
                None => {
                    unique.insert(*c, true);
                    len = len + 1;
                    right = right + 1;
                },
                Some(_) => {
                    let c2 = s.get(left).unwrap();
                    unique.remove(&c2);
                    left = left + 1;
                    len = len - 1;
                }
            }
            if len > max {
                max = len
            }
        }
        max
    }
}

fn main() {
	assert_eq!(0, Solution::length_of_longest_substring("".to_string()));
	assert_eq!(1, Solution::length_of_longest_substring("a".to_string()));
	assert_eq!(1, Solution::length_of_longest_substring("aa".to_string()));
	assert_eq!(2, Solution::length_of_longest_substring("ab".to_string()));
	assert_eq!(3, Solution::length_of_longest_substring("abc".to_string()));
	assert_eq!(4, Solution::length_of_longest_substring("abcd".to_string()));
	assert_eq!(3, Solution::length_of_longest_substring("abcabcbb".to_string()));
	assert_eq!(1, Solution::length_of_longest_substring("bbbbb".to_string()));
	assert_eq!(3, Solution::length_of_longest_substring("pwwkew".to_string()));
	assert_eq!(5, Solution::length_of_longest_substring("anviaj".to_string()));
}
