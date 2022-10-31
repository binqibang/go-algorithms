package slidingwindow

import "testing"

type testPair struct {
	nums   []int
	k      int
	answer float64
}

func TestFindMaxAverage(t *testing.T) {
	nums := []int{1, 12, -5, -6, 50, 3, 47, 8, -9, 45, 33, -90}
	k := 4
	if FindMaxAverage(nums, k) != 27 {
		t.Error("Excepted answer is 27!")
	}
}

func TestTableFindMaxAverage(t *testing.T) {
	var tests []testPair
	pair := testPair{[]int{10, -9, 22, 32, -3, 20}, 3, 17}
	pair1 := testPair{[]int{1, 1}, 2, 1}
	tests = append(tests, pair)
	tests = append(tests, pair1)
	for _, test := range tests {
		output := FindMaxAverage(test.nums, test.k)
		if output != test.answer {
			t.Error("Test failed: ({}, {}) inputted, {} expected, received: {}.",
				test.nums, test.k, test.answer, output)
		} else {
			t.Logf("Teat passed: (%v, %d) inputted, %.2f received.", test.nums, test.k, output)
		}
	}
}
