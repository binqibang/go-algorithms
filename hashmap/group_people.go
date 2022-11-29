package hashmap

func GroupThePeople(groupSizes []int) [][]int {
	var ans [][]int
	groups := map[int][]int{}
	for i, size := range groupSizes {
		groups[size] = append(groups[size], i)
	}
	for size, people := range groups {
		for i := 0; i < len(people); i += size {
			ans = append(ans, people[i:i+size])
		}
	}
	return ans
}
