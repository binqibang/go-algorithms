package string

// ReverseWords
// @summary: LeetCode #151 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/9/16 11:17
func ReverseWords(s string) string {
	var ans []byte
	var start, end int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		}
		end = i
		for i >= 0 && s[i] != ' ' {
			i--
		}
		start = i + 1
		for j := start; j <= end; j++ {
			ans = append(ans, s[j])
		}
		ans = append(ans, ' ')
	}
	ans = ans[0 : len(ans)-1]
	return string(ans)
}
