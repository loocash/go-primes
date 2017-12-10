package primes

import "testing"
import "fmt"

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

func TestGetPrimes(t *testing.T) {
	firstTenPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	result := GetPrimes(30)

	assertSlice(t, firstTenPrimes, result)
}

func ExampleGetPrimes() {
	primesUpTo19 := GetPrimes(19)
	fmt.Println(primesUpTo19)
	// Output:
	// [2 3 5 7 11 13 17 19]
}

func BenchmarkGetPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimes(1000)
	}
}

func TestSieve(t *testing.T) {
	sieveOf10 := []bool{false, false, true, true, false, true, false, true, false, false, false}
	result := Sieve(10)
	assertSlice(t, sieveOf10, result)
}

func ExampleSieve() {
	isPrime := Sieve(10)
	fmt.Println(isPrime[10], isPrime[3])
	// Output:
	// false true
}

func BenchmarkSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sieve(100000)
	}
}

func TestFactorize(t *testing.T) {
	var tests = []struct {
		n       int
		factors []int
	}{
		{55, []int{5, 11}},
		{83, []int{83}},
		{1, []int{}},
		{24, []int{2, 2, 2, 3}},
	}

	primes := GetPrimes(100)

	for _, test := range tests {
		got := Factorize(test.n, primes)
		assertSlice(t, test.factors, got)
	}
}

func ExampleFactorize() {
	num := 24
	factors := Factorize(num, GetPrimes(num))
	fmt.Println(factors)
	// Output:
	// [2 2 2 3]
}

func BenchmarkFactorize(b *testing.B) {
	num := 1000
	primes := GetPrimes(num)
	for i := 0; i < b.N; i++ {
		Factorize(num, primes)
	}
}

func TestTotient(t *testing.T) {
	totientOf10 := []int{0, 1, 1, 2, 2, 4, 2, 6, 4, 6, 4}
	result := Totient(10)
	assertSlice(t, totientOf10, result)
}

func ExampleTotient() {
	fmt.Println(Totient(10))
	// Output:
	// [0 1 1 2 2 4 2 6 4 6 4]
}

func BenchmarkTotient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Totient(1000)
	}
}
