/**
 * Definition of Interval:
 * #[derive(Debug, Clone)]
 * pub struct Interval {
 *     pub start: i32,
 *     pub end: i32,
 * }
 *
 * impl Interval {
 *     pub fn new(start: i32, end: i32) -> Self {
 *         Interval { start, end }
 *     }
 * }
 */

impl Solution {
    pub fn can_attend_meetings(intervals: Vec<Interval>) -> bool {
        let mut intervals = intervals;
        
        if intervals.len() == 0 {
            return true;
        }

        intervals.sort_by_key(|interval| interval.start);

        for i in 1..intervals.len() {
            if intervals[i].start < intervals[i-1].end {
                return false;
            }
        }
        true
    }
}
