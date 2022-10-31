package string

import "strings"

// IsPrefixOfWord
// @summary: LeetCode #1455 (Easy)
// @author : binqibang
// @created: 2022/8/21 20:13
func IsPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], searchWord) {
			return i + 1
		}
	}
	return -1
}
