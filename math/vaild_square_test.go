package math

import "testing"

func TestValidSquare(t *testing.T) {
	p1 := []int{0, 0}
	p2 := []int{1, 1}
	p3 := []int{1, 0}
	p4 := []int{0, 1}
	output := ValidSquare(p1, p2, p3, p4)
	if output != true {
		t.Errorf("Test failed: (%v, %v, %v, %v) inputted, %t expected, received: %t.",
			p1, p2, p3, p4, true, output)
	} else {
		t.Logf("Teat passed: (%v, %v, %v, %v) inputted, %t received,",
			p1, p2, p3, p4, output)
	}
}
