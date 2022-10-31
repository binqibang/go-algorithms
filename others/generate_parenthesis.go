package others

// GenerateParenthesis
// @summary: LeetCode #22 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/10/13 21:22
func GenerateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	var void struct{}
	set := make(map[string]struct{})
	for _, str := range GenerateParenthesis(n - 1) {
		for i := 0; i <= len(str)/2; i++ {
			set[str[:i]+"()"+str[i:]] = void
		}
	}
	var pars []string
	for k := range set {
		pars = append(pars, k)
	}
	return pars
}
