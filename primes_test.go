package primes

import "testing"
import "fmt"

func equalsInt(a, b []int) bool {
	lena, lenb := len(a), len(b)
	if lena != lenb {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalsBool(a, b []bool) bool {
	lena, lenb := len(a), len(b)
	if lena != lenb {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestGetPrimes(t *testing.T) {
	firstTenPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	result := GetPrimes(30)

	if !equalsInt(firstTenPrimes, result) {
		t.Errorf("First 10 prime numbers are %v and not %v", firstTenPrimes, result)
	}

}

func ExampleGetPrimes() {
	primesUpTo19 := GetPrimes(19)
	fmt.Println(primesUpTo19)
	// Output:
	// [2 3 5 7 11 13 17 19]
}

func TestSieve(t *testing.T) {
	sieveOf10 := []bool{false, false, true, true, false, true, false, true, false, false, false}
	result := Sieve(10)
	if !equalsBool(sieveOf10, result) {
		t.Errorf("Sieve of 10 is %v and not %v", sieveOf10, result)
	}
}

func ExampleSieve() {
	isPrime := Sieve(10)
	fmt.Println(isPrime[10], isPrime[3])
	// Output:
	// false true
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
		if factors := Factorize(test.n, primes); !equalsInt(factors, test.factors) {
			t.Errorf("%d factorized is %v and not %v\n", test.n, test.factors, factors)
		}
	}
}

func ExampleFactorize() {
	num := 24
	factors := Factorize(num, GetPrimes(num))
	fmt.Println(factors)
	// Output:
	// [2 2 2 3]
}

func TestTotient(t *testing.T) {
	totientOf10 := []int{0, 1, 1, 2, 2, 4, 2, 6, 4, 6, 4}
	result := Totient(10)
	if !equalsInt(totientOf10, result) {
		t.Errorf("Totient of 10 should be %v instead of %v\n", totientOf10, result)
	}
}

func ExampleTotient() {
	fmt.Println(Totient(10))
	// Output:
	// [0 1 1 2 2 4 2 6 4 6 4]
}
