struct Solution;

impl Solution {
    pub fn int_to_roman(num: i32) -> String {
        let m = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
        let s = [
            "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I",
        ];
        let mut n = num;
        let mut buf = vec![];
        for i in 0..m.len() {
            let r = n / m[i];
            n = n % m[i];
            for _ in 0..r {
                buf.push(s[i]);
            }
        }
        buf.into_iter().collect()
    }
}

fn main() {
    for i in 1..30 {
        println!("{}", Solution::int_to_roman(i));
    }
}
