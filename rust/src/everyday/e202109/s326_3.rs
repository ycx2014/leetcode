struct Solution;
impl Solution {
    pub fn is_power_of_three(n: i32) -> bool {
        return n > 0 && 1162261467 % n == 0;
    }
}
