package primes

import "testing"

func TestUniq(t *testing.T) {
	var tests = []struct {
		before, after []int
	}{
		{[]int{1, 1, 1, 1, 1}, []int{1}},
		{[]int{1, 1, 2, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 1, 1, 2}, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		got := uniq(test.before)
		assertSlice(t, test.after, got)
	}
}
