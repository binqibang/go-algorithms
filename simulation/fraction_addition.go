package simulation

import (
	"fmt"
	"unicode"
)

// FractionAddition
// @summary: LeetCode #592 (Medium)
// @author : binqibang
// @create : 2022-07-27 11:19
func FractionAddition(expression string) string {
	denominator, numerator := 0, 1
	idx := 0
	for idx < len(expression) {
		// 读取分子
		// sign: 1 or -1
		denominator1, sign := 0, 1
		if expression[idx] == '-' || expression[idx] == '+' {
			if expression[idx] == '-' {
				sign = -1
			}
			idx++
		}
		for idx < len(expression) && unicode.IsDigit(rune(expression[idx])) {
			denominator1 = denominator1*10 + int(expression[idx]-'0')
			idx++
		}
		denominator1 *= sign

		// 跳过 '/'
		idx++

		// 读取分母
		numerator1 := 0
		for idx < len(expression) && unicode.IsDigit(rune(expression[idx])) {
			numerator1 = numerator1*10 + int(expression[idx]-'0')
			idx++
		}

		// 计算
		denominator = denominator1*numerator + denominator*numerator1
		numerator *= numerator1
	}
	if denominator == 0 {
		return "0/1"
	}
	g := gcd(abs(denominator), numerator)
	return fmt.Sprintf("%d/%d", denominator/g, numerator/g)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func gcd(a, b int) int {
	reminder := a % b
	for reminder != 0 {
		a, b = b, reminder
		reminder = a % b
	}
	return b
}
