package graph

// CanFinish
// @summary: LeetCode #207 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/9/5 11:13
func CanFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges    = make([][]int, numCourses)
		indegree = make([]int, numCourses)
	)

	for _, edge := range prerequisites {
		edges[edge[1]] = append(edges[edge[1]], edge[0])
		indegree[edge[0]]++
	}

	var queue []int
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	visited := 0
	for len(queue) != 0 {
		visited++
		u := queue[0]
		queue = queue[1:]
		for _, v := range edges[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return visited == numCourses
}
