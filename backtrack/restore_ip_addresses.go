package backtrack

import "strconv"

// RestoreIpAddresses
// @summary: LeetCode #93 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/9/29 10:18

const SegCount = 4

var (
	ans      []string
	segments []int
)

func RestoreIpAddresses(s string) []string {
	// reinitialize
	segments = make([]int, SegCount)
	ans = []string{}
	if len(s) < 4 || len(s) > 12 {
		return ans
	}
	dfs(s, 0, 0)
	return ans
}

func dfs(s string, segId, segIndex int) {
	// found 4 IP segments and visited all chars
	if segId == SegCount {
		if segIndex == len(s) {
			ip := ""
			for i := 0; i < SegCount; i++ {
				ip += strconv.Itoa(segments[i])
				if i != SegCount-1 {
					ip += "."
				}
			}
			ans = append(ans, ip)
		}
		return
	}

	// visited all chars but not found 4 IP segments
	if segIndex == len(s) {
		return
	}

	// if visit '0', then this segment will be 0
	if s[segIndex] == '0' {
		segments[segId] = 0
		dfs(s, segId+1, segIndex+1)
	}

	// common
	seg := 0
	for i := segIndex; i < len(s); i++ {
		seg = seg*10 + int(s[i]-'0')
		if seg > 0 && seg <= 255 {
			segments[segId] = seg
			dfs(s, segId+1, i+1)
		} else {
			break
		}
	}
}
