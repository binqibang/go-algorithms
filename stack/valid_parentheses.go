package stack

// IsValid
// @summary: LeetCode #20 (Easy); CodeTop
// @author : binqibang
// @created: 2022/10/18 11:47
func IsValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var stack []byte
	for i := 0; i < n; i++ {
		ch := s[i]
		if _, ok := pairs[ch]; ok { // left part, directly push
			stack = append(stack, ch)
		} else { // right part, compare with pop one
			if len(stack) == 0 || pairs[stack[len(stack)-1]] != ch {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
