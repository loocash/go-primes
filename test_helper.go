package primes

import (
	"fmt"
	"testing"
)

func equalsSlice(a, b interface{}) bool {
	astr := fmt.Sprintf("%v", a)
	bstr := fmt.Sprintf("%v", b)
	return astr == bstr
}

func assertSlice(t *testing.T, want, got interface{}) {
	if !equalsSlice(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
