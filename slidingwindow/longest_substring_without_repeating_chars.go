package slidingwindow

// LengthOfLongestSubstring
// @summary: LeetCode #3 (Medium), CodeTop MS
// @author : binqibang
// @created: 2022/8/24 10:29
func LengthOfLongestSubstring(s string) int {
	indices := map[byte]int{}
	l := 0
	ans := 0
	for r := 0; r < len(s); r++ {
		if idx, ok := indices[s[r]]; ok {
			l = max(l, idx+1)
		}
		ans = max(ans, r-l+1)
		indices[s[r]] = r
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
