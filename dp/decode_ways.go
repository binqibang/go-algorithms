package dp

// NumDecodings
// @summary: LeetCode #91 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/11/2 11:07
func NumDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	var idx int
	for i := 1; i <= n; i++ {
		idx = i - 1
		if s[idx] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[idx-1] != '0' && (10*(s[idx-1]-'0')+s[idx]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
