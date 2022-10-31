package string

import "math"

// MyAtoi
// @summary: LeetCode #8 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/10/25 10:18
func MyAtoi(s string) int {
	var ans int
	idx, n, sign := 0, len(s), 1
	// remove all leading spaces
	for idx < n && s[idx] == ' ' {
		idx++
	}
	// visits '+' or '-'
	if idx < n && (s[idx] == '+' || s[idx] == '-') {
		if s[idx] == '-' {
			sign = -1
		}
		idx++
	}
	// visit digits '0' ~ '9',
	// if visits no-digit char, then finish loop
	for idx < n && '0' <= s[idx] && s[idx] <= '9' {
		digit := int(s[idx] - '0')
		// integer overflow, clamp
		if ans > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		ans = ans*10 + digit
		idx++
	}
	return ans * sign
}
