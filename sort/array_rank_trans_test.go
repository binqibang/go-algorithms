package sort

import "testing"

type testPair struct {
	arr    []int
	answer []int
}

func check(a, b []int) bool {
	// 两个 slice 长度不同，必然不等
	if len(a) != len(b) {
		return false
	}
	// 两个 slice 只有一个为空，必然不等
	if (a == nil) != (b == nil) {
		return false
	}
	// 遍历比较各个元素
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestArrayRankTransform(t *testing.T) {
	// Prepare test data.
	var tests []testPair
	test := testPair{[]int{100, 100, 100}, []int{1, 1, 1}}
	test1 := testPair{[]int{37, 12, 28, 9, 100, 56, 80, 5, 12}, []int{5, 3, 4, 2, 8, 6, 7, 1, 3}}
	tests = append(tests, test)
	tests = append(tests, test1)

	// Tests.
	for _, test := range tests {
		// 将 test.arr 复制到 arr，且不影响原切片
		arr := append([]int{}, test.arr...)
		output := ArrayRankTransform(test.arr)
		if !check(output, test.answer) {
			t.Errorf("Test failed: %+v inputted, %+v expected, received: %+v.",
				arr, test.answer, output)
		} else {
			t.Logf("Teat passed: %+v inputted, %+v received.", arr, output)
		}
	}

}
