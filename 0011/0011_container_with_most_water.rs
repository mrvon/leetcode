struct Solution;

impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        if height.len() < 2 {
            return 0;
        }
        let mut max_area: i32 = 0;
        let mut l_range: i32 = 0;
        let mut r_range: i32 = height.len() as i32 - 1;
        while l_range < r_range {
            let m = i32::min(height[l_range as usize], height[r_range as usize]) * (r_range - l_range);
            if m > max_area {
                max_area = m;
            }
            if height[l_range as usize] <= height[r_range as usize] {
                l_range += 1;
            } else {
                r_range -= 1;
            }
        }
        return max_area;
    }
}

fn main() {
	assert_eq!(8, Solution::max_area(vec![ 1, 2, 2, 3, 4, 6 ]));
}
