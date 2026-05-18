use std::collections::HashSet;

impl Solution {
    pub fn has_duplicate(nums: Vec<i32>) -> bool {
       let mut seen = HashSet::new();

       for n in nums {
        if seen.contains(&n) {
            return true
        }

        seen.insert(n);
       }

       false
    }
}
