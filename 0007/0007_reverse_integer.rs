struct Solution;

impl Solution {
    pub fn reverse(x: i32) -> i32 {
        x.signum() * x
			.abs()
			.to_string()
			.chars()
			.rev()
			.collect::<String>()
			.parse::<i32>()
			.unwrap_or(0)
    }
}

fn main() {
	assert_eq!(1, Solution::reverse(1));
	assert_eq!(1, Solution::reverse(10));
	assert_eq!(1, Solution::reverse(100));
	assert_eq!(123, Solution::reverse(321));
	assert_eq!(-123, Solution::reverse(-321));
	assert_eq!(0, Solution::reverse(2147483647));
	assert_eq!(0, Solution::reverse(-2147483647));
}
