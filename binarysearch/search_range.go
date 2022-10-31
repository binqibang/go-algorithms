package binarysearch

// SearchRange
// @summary: LeetCode #34 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/10/10 15:40
func SearchRange(nums []int, target int) []int {
	startIdx := binarySearch(nums, target, true)
	endIdx := binarySearch(nums, target, false) - 1
	if startIdx <= endIdx && endIdx < len(nums) && nums[startIdx] == target &&
		nums[endIdx] == target {
		return []int{startIdx, endIdx}
	}
	return []int{-1, -1}
}

func binarySearch(nums []int, target int, equal bool) int {
	left, right, res := 0, len(nums)-1, len(nums)
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target || (equal && nums[mid] >= target) {
			right = mid - 1
			res = mid
		} else {
			left = mid + 1
		}
	}
	return res
}
