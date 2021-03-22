package algos

import (
	"testing"
	"fmt"
)

func TestBinarySearchTableDriven(t *testing.T) {
	var tests = []struct {
		a [5]int
		target int
		want int
	}{
		{[5]int{1, 2, 3, 4, 5}, 3, 2},
		{[5]int{1, 2, 3, 4, 5}, 10, -1},
		{[5]int{1, 2, 3, 4, 5}, 0, -1},
	}
	for _, test := range tests {
		testname := fmt.Sprintf("search for %d", test.target)
		t.Run(testname, func(t *testing.T) {
			res := BinarySearch(test.a, test.target)
			if res != test.want {
				t.Errorf("got %d, want %d", res, test.want)
			}
		})
	}
}