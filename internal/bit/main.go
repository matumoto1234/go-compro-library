package bit

func BitCeil(n uint64) uint64 {
	x := uint64(1)
	for x < n {
		x <<= 1
	}
	return x
}

func CountRightZero(n uint64) int {
	if n == 0 {
		return 64
	}
	c := 0
	for n&1 == 0 {
		n >>= 1
		c++
	}
	return c
}