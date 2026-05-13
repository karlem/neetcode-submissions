use std::collections::HashMap;

impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        if s.len() != t.len() {
            return false;
        }

        let mut seen = HashMap::new();

        for c in s.chars() {
            *seen.entry(c).or_insert(0) += 1;
        }
        
        for c in t.chars() {
            *seen.entry(c).or_insert(0) -= 1;
        }

        seen.values().all(|&x| x==0)
    }
}
