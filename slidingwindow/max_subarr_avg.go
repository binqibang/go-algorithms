// Package slidingwindow
// @summary: LeetCode questions about "sliding window"
// @author : binqibang
// @create : 2022-07-22 11:43
package slidingwindow

import "math"

// FindMaxAverage LeetCode #634 (Easy)
func FindMaxAverage(nums []int, k int) float64 {
	var sum = 0
	var ans = -10e4
	for _, num := range nums[:k] {
		sum += num
	}
	l := 0
	r := l + k - 1
	for r < len(nums) {
		ans = math.Max(ans, float64(sum)/float64(k))
		if r == len(nums)-1 {
			break
		}
		r++
		sum += nums[r]
		sum -= nums[l]
		l++
	}
	return ans
}
