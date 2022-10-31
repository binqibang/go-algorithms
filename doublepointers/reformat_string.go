// Package doublepointers @summary: LeetCode questions about "double pointers"
// @author : binqibang
// @created: 2022/8/11 21:41
package doublepointers

import "unicode"

func Reformat(s string) string {
	numDigit := 0
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			numDigit++
		}
	}
	numAlpha := len(s) - numDigit
	if abs(numAlpha-numDigit) > 1 {
		return ""
	}
	ans := []byte(s)
	// 数量少的字符类型
	flag := numDigit > numAlpha
	for i, j := 0, 1; i < len(ans); i += 2 {
		// 偶数位上是数量少的类型
		if unicode.IsDigit(rune(ans[i])) != flag {
			for unicode.IsDigit(rune(ans[j])) != flag {
				// 从前往后隔位寻找
				j += 2
			}
			ans[i], ans[j] = ans[j], ans[i]
		}
	}
	return string(ans)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
