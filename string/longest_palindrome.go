package string

// LongestPalindrome
// @summary: LeetCode #5 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/9/23 11:30
func LongestPalindrome(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}
	var left, right int
	curLen, maxLen := 1, 0
	start := 0
	for i := 0; i < n; i++ {
		left, right = i-1, i+1
		// left diffusion to find same char
		for left >= 0 && s[left] == s[i] {
			left--
			curLen++
		}
		// right diffusion to find same char
		for right < n && s[right] == s[i] {
			right++
			curLen++
		}
		// left & right diffusion until `left char != right char`
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
			curLen += 2
		}
		if curLen > maxLen {
			maxLen = curLen
			start = left + 1
		}
		curLen = 1
	}
	return s[start : start+maxLen]
}
