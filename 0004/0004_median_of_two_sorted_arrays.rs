struct Solution;

impl Solution {
    pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
        let mut nums3: Vec<i32> = vec![];
        nums3.extend(nums1.iter());
        nums3.extend(nums2.iter());
        nums3.sort();
        if nums3.len() % 2 == 0 {
            (nums3[nums3.len() / 2 - 1] + nums3[nums3.len() / 2]) as f64 / 2.0
        } else {
            nums3[nums3.len() / 2] as f64
        }
    }
}

fn main() {
    assert_eq!(1.0, Solution::find_median_sorted_arrays(vec![1], vec![]));
    assert_eq!(2.0, Solution::find_median_sorted_arrays(vec![], vec![2]));
    assert_eq!(1.5, Solution::find_median_sorted_arrays(vec![1], vec![2]));
    assert_eq!(
        2.0,
        Solution::find_median_sorted_arrays(vec![1, 3], vec![2])
    );
    assert_eq!(
        2.5,
        Solution::find_median_sorted_arrays(vec![1, 2], vec![3, 4])
    );
    assert_eq!(
        2.0,
        Solution::find_median_sorted_arrays(vec![1, 2], vec![1, 2, 3])
    );
    assert_eq!(
        3.5,
        Solution::find_median_sorted_arrays(vec![3, 4, 5], vec![1, 2, 7])
    );
    assert_eq!(
        5.5,
        Solution::find_median_sorted_arrays(vec![1, 3, 5, 7, 9], vec![2, 4, 6, 8, 10])
    );
}
