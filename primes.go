// Package primes generates prime numbers
package primes

import "math"

// GetPrimes returns a slice of all primes less than or equal to n in ascending order
func GetPrimes(n int) []int {
	var p []int
	sieve := Sieve(n)
	for i, isPrime := range sieve {
		if isPrime {
			p = append(p, i)
		}
	}
	return p
}

// Sieve returns a slice of booleans telling about which number is prime
func Sieve(n int) []bool {

	var p []bool

	p = make([]bool, n+1)

	for i := 2; i <= n; i++ {
		p[i] = true
	}

	upperBound := int(math.Ceil(math.Sqrt(float64(n))))

	for i := 2; i < upperBound; i++ {
		for j := i * i; j <= n; j += i {
			p[j] = false
		}
	}

	return p
}

// Factorize returns an array of prime factors of number n
func Factorize(n int, primes []int) []int {
	var factors []int
	for _, prime := range primes {
		if prime > n {
			break
		}
		for n%prime == 0 {
			factors = append(factors, prime)
			n /= prime
		}
	}
	return factors
}

// Totient function counts the positive integers up to a given integer n that are relatively prime to n
func Totient(n int) []int {
	var t []int
	//primes := GetPrimes(n)

	t = make([]int, n+1)
	for i := 1; i <= n; i++ {
		t[i] = 1
	}

	for i := 2; i <= n; i++ {
		//factors := Factorize(i, primes)
		// TODO
	}

	return t
}
