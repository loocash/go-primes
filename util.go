package primes

// uniq returns a slice with adjacent duplicates removed
func uniq(a []int) []int {
	last := 0
	for i := range a {
		if a[i] != a[last] {
			last++
			a[last] = a[i]
		}
	}
	return a[:last+1]
}
