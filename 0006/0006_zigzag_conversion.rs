struct Solution;

impl Solution {
    pub fn convert(s: String, num_rows: i32) -> String {
        let mut rows: Vec<Vec<char>> = vec![vec![]; num_rows as usize];
        let mut idx: i32 = 0;
        let mut dir: i32 = 1;
        for c in s.chars() {
            rows[idx as usize].push(c);
            if dir == 1 && idx >= num_rows-1 {
                dir = -1;
            } else if dir == -1 && idx <= 0 {
                dir = 1;
            }
            if num_rows > 1 {
                idx += dir;
            }
        }
        let mut result: Vec<char> = vec![];
        for v in rows {
            result.extend(v.iter());
        }
        result.iter().collect()
    }
}

fn main() {
	assert_eq!("PAYPALISHIRING".to_string(), Solution::convert("PAYPALISHIRING".to_string(), 1));
	assert_eq!("PYAIHRNAPLSIIG".to_string(), Solution::convert("PAYPALISHIRING".to_string(), 2));
	assert_eq!("PAHNAPLSIIGYIR".to_string(), Solution::convert("PAYPALISHIRING".to_string(), 3));
	assert_eq!("PINALSIGYAHRPI".to_string(), Solution::convert("PAYPALISHIRING".to_string(), 4));
}
