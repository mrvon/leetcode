struct Solution;

impl Solution {
    pub fn get_value(ch: char) -> i32 {
        match ch {
            'I' => 1,
            'V' => 5,
            'X' => 10,
            'L' => 50,
            'C' => 100,
            'D' => 500,
            'M' => 1000,
            _ => 0,
        }
    }

    pub fn roman_to_int(s: String) -> i32 {
        let mut sum = 0;
        let mut last = 0;
        for c in s.chars().rev() {
            let curr = Self::get_value(c);
            if curr >= last {
                sum += curr;
            } else {
                sum -= curr;
            }
            last = curr;
        }
        sum
    }
}

fn main() {
    assert_eq!(Solution::roman_to_int("III".to_string()), 3);
    assert_eq!(Solution::roman_to_int("IV".to_string()), 4);
}
