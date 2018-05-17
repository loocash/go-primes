// Package primes generates prime numbers
package primes

import (
	"sort"
)

// UpTo returns a slice of all primes less than or equal to n in ascending order
func UpTo(n int) []int {
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

	isPrime := make([]bool, n+1)

	// initially all numbers from 2 to n are assumed to be prime
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// for every prime number p
	for p := 2; p*p < n; p++ {
		if !isPrime[p] {
			continue
		}
		// remove all numbers j = p^2 + k*p, where k >= 0
		for j := p * p; j <= n; j += p {
			isPrime[j] = false
		}
	}

	return isPrime
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
	primes := UpTo(n)

	t = make([]int, n+1)
	for i := 1; i <= n; i++ {
		t[i] = 1
	}

	for _, prime := range primes {
		for pow := prime; pow <= n; pow *= prime {
			for j := pow; j <= n; j += pow {
				if pow == prime {
					t[j] *= prime - 1
				} else {
					t[j] *= prime
				}
			}
		}
	}

	/*
		for i := 2; i <= n; i++ {
			factors := Factorize(i, primes)
			for j, factor := range factors {
				if j == 0 || factors[j-1] != factors[j] {
					t[i] *= factor - 1
				} else {
					t[i] *= factor
				}
			}
		}
	*/

	return t
}

// Divisors returns every divisor of n as an array in ascending order
func Divisors(n int, primes []int) []int {
	factors := Factorize(n, primes)
	flen := len(factors)
	power := 1 << uint(flen)
	divs := make([]int, 0, power)
	for i := 0; i < power; i++ {
		n := 1
		for fi, f := range factors {
			if i&(1<<uint(fi)) > 0 {
				n *= f
			}
		}
		divs = append(divs, n)
	}

	sort.Ints(divs)

	return uniq(divs)
}
