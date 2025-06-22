package math

import "golang.org/x/exp/constraints"

// a^-1 (mod m)
func ModInv[T constraints.Integer](a, m T) T {
	// a*x + mod*y = 1
	_, x, _ := ExtGCD(a, m)
	if x < 0 {
		x += m
	}
	return x % m
}
