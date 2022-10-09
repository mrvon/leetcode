struct Solution;

impl Solution {
    pub fn generate_parenthesis(n: i32) -> Vec<String> {
        let mut results: Vec<String> = vec![];
        let mut buf: Vec<char> = vec![];
        for _ in 0..n {
            buf.push('(');
        }
        for _ in 0..n {
            buf.push(')');
        }
        // this is must be well-formed
        results.push(buf.iter().collect());
        while Self::next_permutation(&mut buf) {
            if Self::is_wellformed(&buf) {
                results.push(buf.iter().collect());
            }
        }
        results
    }

    fn next_permutation(buf: &mut Vec<char>) -> bool {
        let n = buf.len() - 1;
        let mut k2: Option<usize> = None;
        for i in (0..=n-1).rev() {
            if buf[i] < buf[i + 1] {
                k2 = Some(i);
                break;
            }
        }
        // last permutation
        if k2.is_none() {
            return false;
        }
        let k = k2.unwrap();
        let mut l = n;
        for i in (k + 1..=n).rev() {
            if buf[k] < buf[i] {
                l = i;
                break;
            }
        }
        buf.swap(k, l);
        let mut i = k + 1;
        let mut j = n;
        while i < j {
            buf.swap(i, j);
            i += 1;
            j -= 1;
        }
        true
    }

    fn is_wellformed(buf: &Vec<char>) -> bool {
        let mut c = 0;
        let mut e = false;
        buf.iter().for_each(|x| {
            if x == &'(' {
                c += 1;
            } else {
                c -= 1;
                if c < 0 {
                    e = true
                }
            }
        });
        !e && c == 0
    }
}

fn main() {
    assert_eq!(vec!["()"], Solution::generate_parenthesis(1));
    assert_eq!(vec!["(())", "()()"], Solution::generate_parenthesis(2));
    assert_eq!(
        vec!["((()))", "(()())", "(())()", "()(())", "()()()"],
        Solution::generate_parenthesis(3)
    );
}
