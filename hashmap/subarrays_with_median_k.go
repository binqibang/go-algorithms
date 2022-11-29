package hashmap

// CountSubarrays
// @summary: LeetCode #2488 (Hard); CodeTop MS
// @author : binqibang
// @created: 2022/11/29 11:24
func CountSubarrays(nums []int, k int) int {
	pos, n := 0, len(nums)
	for nums[pos] != k {
		pos++
	}
	rightCount := map[int]int{
		0: 1,
	}
	for i, sum := pos+1, 0; i < n; i++ {
		if nums[i] > k {
			sum++
		} else {
			sum--
		}
		rightCount[sum]++
	}
	ans := 0
	ans += rightCount[0] + rightCount[1]
	for i, sum := pos-1, 0; i >= 0; i-- {
		if nums[i] > k {
			sum++
		} else {
			sum--
		}
		ans += rightCount[-1*sum] + rightCount[1-sum]
	}
	return ans
}
