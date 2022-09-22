struct Solution;

impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        match strs.is_empty() {
            true => "".to_string(),
            _ => strs.iter().skip(1).fold(strs[0].clone(), |acc, x| {
                acc.chars()
                    .zip(x.chars())
                    .take_while(|(x, y)| x == y)
                    .map(|(x, _)| x)
                    .collect()
            }),
        }
    }
}

fn main() {
    assert_eq!(
        "",
        Solution::longest_common_prefix(vec!(
            "hel".to_string(),
            "hell".to_string(),
            "aello".to_string()
        ))
    );

    assert_eq!(
        "hello",
        Solution::longest_common_prefix(vec!(
            "hello".to_string(),
            "hello world".to_string(),
            "hello rust".to_string()
        ))
    );
}
