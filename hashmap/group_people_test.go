package hashmap

import "testing"

func TestGroupThePeople(t *testing.T) {
	test := []int{2, 1, 3, 3, 3, 2}
	t.Log(GroupThePeople(test))

	test1 := []int{3, 3, 3, 3, 3, 1, 3}
	t.Log(GroupThePeople(test1))
}
