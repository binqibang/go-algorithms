package graph

import "math"

// NetworkDelayTime
// @summary: LeetCode #743 (Medium);
// @author : binqibang
// @created: 2022/12/8 16:10
func NetworkDelayTime(times [][]int, n int, k int) int {
	const INF = math.MaxInt64 / 2
	// initialize adjacency matrix
	var graph = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph[i][j] = INF
		}
	}
	for _, t := range times {
		u, v, w := t[0]-1, t[1]-1, t[2]
		graph[u][v] = w
	}
	// dist(k -> k) = 0, others = INF
	var dist = make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	dist[k-1] = 0
	// if flag[i] is true, vertex `i` is processed
	var flag = make([]bool, n)
	for i := 1; i < n; i++ {
		u, minDist := 0, INF
		for v := 0; v < n; v++ {
			if !flag[v] && dist[v] < minDist {
				minDist = dist[v]
				u = v
			}
		}
		flag[u] = true
		for v := 0; v < n; v++ {
			dist[v] = min(dist[v], dist[u]+graph[u][v])
		}
	}
	ans := 0
	for _, d := range dist {
		if d == INF {
			return -1
		}
		ans = max(ans, d)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
