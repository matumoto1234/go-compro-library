package math

import "github.com/matumoto1234/go-compro-library/math/constraints"

// a^-1 (mod m)
func ModInv[T constraints.Integer](a, m T) T {
	// a*x + mod*y = 1
	_, x, _ := ExtGCD(a, m)
	if x < 0 {
		x += m
	}
	return x % m
}
