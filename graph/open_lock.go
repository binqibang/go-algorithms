package graph

// OpenLock
// @summary: LeetCode #752 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/11/9 15:30
func OpenLock(deadends []string, target string) int {
	const start = "0000"

	if target == start {
		return 0
	}

	dead := map[string]bool{}
	for _, end := range deadends {
		dead[end] = true
	}
	if dead[start] {
		return -1
	}

	tried := map[string]bool{start: true}
	queue := []string{start}
	step := 0

	for len(queue) != 0 {
		step++
		n := len(queue)
		for i := 0; i < n; i++ {
			code := queue[0]
			queue = queue[1:]
			for _, next := range GetNextCodes(code) {
				if !dead[next] && !tried[next] {
					if next == target {
						return step
					}
					tried[next] = true
					queue = append(queue, next)
				}
			}
		}
	}

	return -1
}

func GetNextCodes(code string) (res []string) {
	// string -> []byte
	bytes := []byte(code)
	for i, b := range bytes {
		bytes[i] = b - 1
		if bytes[i] < '0' {
			bytes[i] = '9'
		}
		// []byte -> string
		res = append(res, string(bytes))

		bytes[i] = b + 1
		if bytes[i] > '9' {
			bytes[i] = '0'
		}
		res = append(res, string(bytes))

		bytes[i] = b
	}
	return
}
