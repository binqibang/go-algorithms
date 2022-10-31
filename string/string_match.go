package string

import (
	"sort"
	"strings"
)

// MatchString
// @summary: LeetCode #1408 (Easy)
// @author : binqibang
// @created: 2022/8/6 11:24
func MatchString(words []string) []string {
	var ans []string
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], words[i]) {
				ans = append(ans, words[i])
				break
			}
		}
	}
	return ans
}
