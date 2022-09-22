struct Solution;

impl Solution {
    pub fn three_sum(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut result: Vec<Vec<i32>> = vec![];
        let n = nums.len();
        nums.sort();
        let mut i = 0;
        while i < n {
            if nums[i] > 0 {
                break;
            }
            let mut j = i + 1;
            while j < n {
                if nums[i] + nums[j] > 0 {
                    break;
                }
                let mut l = j + 1; // left range
                let mut r = n - 1; // right range
                                   // binary search
                let target = -nums[i] - nums[j];
                while l <= r {
                    let m = l + (r - l) / 2;
                    if target == nums[m] {
                        result.push(vec![nums[i], nums[j], nums[m]]);
                        break;
                    } else if target < nums[m] {
                        r = m - 1;
                    } else {
                        l = m + 1;
                    }
                }
                // next j
                j += 1;
                while j < n && nums[j - 1] == nums[j] {
                    j += 1;
                }
            }
            // next i
            i += 1;
            // skip same element
            while i < n && nums[i - 1] == nums[i] {
                i += 1;
            }
        }
        result
    }
}

fn main() {
    assert_eq!(vec![] as Vec<Vec<i32>>, Solution::three_sum(vec![]));
    assert_eq!(vec![] as Vec<Vec<i32>>, Solution::three_sum(vec![0]));
    assert_eq!(vec![vec![0, 0, 0]] as Vec<Vec<i32>>, Solution::three_sum(vec![0, 0, 0]));
    assert_eq!(
        vec![vec![-1, -1, 2], vec![-1, 0, 1]] as Vec<Vec<i32>>,
        Solution::three_sum(vec![-1, 0, 1, 2, -1, -4])
    );
}
