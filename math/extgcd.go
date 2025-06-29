package math

import "github.com/matumoto1234/go-compro-library/math/constraints"

// return (gcd, x, y)
func ExtGCD[T constraints.Integer](a, b T) (T, T, T) {
	return extGCD(a, b, 0, 0)
}

func extGCD[T constraints.Integer](a, b, x, y T) (T, T, T) {
	if b == 0 {
		// a * 1 + b * 0 = gcd(a, b)
		return a, 1, 0
	}

	q := a / b
	r := a % b

	s := q*x + y
	t := x

	gcd, s, t := extGCD(b, r, s, t)

	return gcd, t, s - q*t
}
