package binarysearch

// Search
// @summary: LeetCode #33 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/10/5 11:10
func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// When the array is split in the `mid`, half of it is always in order,
		// then we binary search in the ordered part.
		if nums[left] <= nums[mid] { // nums[left] ~ nums[mid] are partially ordered
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // nums[mid] ~ nums[right] are partially ordered
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
