package backtrack

// CombinationSum
// @summary: LeetCode #39 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/10/31 15:26
func CombinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	var backtrack func(idx int, t int)
	backtrack = func(idx int, t int) {
		if t == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if candidates[i] <= t {
				path = append(path, candidates[i])
				backtrack(i, t-candidates[i])
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0, target)
	return res
}
