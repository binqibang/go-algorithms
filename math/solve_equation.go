package math

import (
	"strconv"
	"unicode"
)

// SolveEquation
// @summary: LeetCode #640 (Medium)
// @author : binqibang
// @created: 2022/8/10 21:46
func SolveEquation(equation string) string {
	// 系数，值
	factor, val := 0, 0
	idx, n := 0, len(equation)
	// 等式左边默认系数为正
	initialSign := 1
	for idx < n {
		if equation[idx] == '=' {
			// 等式右边默认系数为负
			initialSign = -1
			idx++
		}
		curSign, number := initialSign, 0
		// x 前是否有系数
		valid := false
		// 去除符号
		if equation[idx] == '+' {
			idx++
		} else if equation[idx] == '-' {
			curSign = -curSign
			idx++
		}
		// 遍历到数字
		for idx < n && unicode.IsDigit(rune(equation[idx])) {
			number = number*10 + int(equation[idx]-'0')
			idx++
			valid = true
		}
		if idx < n && equation[idx] == 'x' { // 变量
			if valid { // x
				curSign *= number
			}
			factor += curSign
			idx++
		} else { // 数值
			val += number * curSign
		}
	}
	if factor == 0 {
		if val == 0 {
			return "Infinite solutions"
		}
		return "No solution"
	}
	return "x=" + strconv.Itoa(-val/factor)
}
