struct Solution;

impl Solution {
    pub fn my_atoi(str: String) -> i32 {
        let (n, s) = match str.chars().skip_while(|x| x.is_whitespace()).take(1).next() {
            Some('+') => (1, 1),
            Some(x) if x.is_digit(10) => (0, 1),
            Some('-') => (1, -1),
            _ => return 0,
        };
        let mut res = 0i32;
        let overflow = if s > 0 { std::i32::MAX } else { std::i32::MIN };
        for c in str
            .chars()
            .skip_while(|x| x.is_whitespace())
            .skip(n)
            .take_while(|x| x.is_digit(10))
        {
            let (r, o) = res.overflowing_mul(10);
            if o {
                return overflow;
            }
            let (r, o) = r.overflowing_add(s * (c as i32 - '0' as i32));
            if o {
                return overflow;
            }
            res = r;
        }
        res
    }
}

fn main() {
    assert_eq!(0, Solution::my_atoi("a".to_string()));
    assert_eq!(1, Solution::my_atoi("1".to_string()));
    assert_eq!(1, Solution::my_atoi("1a".to_string()));
    assert_eq!(0, Solution::my_atoi("".to_string()));
    assert_eq!(0, Solution::my_atoi(" ".to_string()));
    assert_eq!(1, Solution::my_atoi(" 1".to_string()));
    assert_eq!(1, Solution::my_atoi(" 1a".to_string()));
    assert_eq!(2147483647, Solution::my_atoi("2147483647".to_string()));
    assert_eq!(2147483647, Solution::my_atoi("2147483648".to_string()));
    assert_eq!(-2147483647, Solution::my_atoi("-2147483647".to_string()));
    assert_eq!(-2147483648, Solution::my_atoi("-2147483648".to_string()));
    assert_eq!(-2147483648, Solution::my_atoi("-2147483649".to_string()));
    assert_eq!(
        2147483647,
        Solution::my_atoi("9223372036854775809".to_string())
    );
    assert_eq!(
        -2147483648,
        Solution::my_atoi("-9223372036854775809".to_string())
    );
    assert_eq!(0, Solution::my_atoi("hello 987".to_string()));
    assert_eq!(0, Solution::my_atoi("+-987".to_string()));
}
