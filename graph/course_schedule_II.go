package graph

// FindOrder
// @summary: LeetCode #210 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/10/7 10:25
func FindOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges    = make([][]int, numCourses)
		indegree = make([]int, numCourses)
		order    []int
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
	for len(queue) != 0 {
		u := queue[0]
		queue = queue[1:]
		order = append(order, u)
		for _, v := range edges[u] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if len(order) == numCourses {
		return order
	}
	return []int{}
}
