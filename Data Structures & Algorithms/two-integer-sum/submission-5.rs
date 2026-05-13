use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut previous = HashMap::new();

        for (i, num) in nums.iter().enumerate() {
            let need = target-num;
            if let Some(&first_index) = previous.get(&need) {
                return vec![first_index as i32, i as i32];
            }

            previous.insert(num, i);
        }

        vec![]
    }
}
