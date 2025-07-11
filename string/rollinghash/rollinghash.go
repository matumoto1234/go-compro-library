package rollinghash

import (
	"math/rand"
	"sync"

	"github.com/matumoto1234/go-compro-library/algorithm/min"
	"github.com/matumoto1234/go-compro-library/math"
)

const mask30 = (1 << 30) - 1
const mask31 = (1 << 31) - 1
const mask61 = (1 << 61) - 1
const Mod = (1 << 61) - 1

var base uint
var once sync.Once

func init() {
	once.Do(func() {
		base = uint(rand.Int63n(Mod))
	})
}

type RollingHash struct {
	Base             uint
	Values           []int
	Inv              []int // Inv[i] := base^-i
	CumulativeHashes []int // CumulativeHashes[i] := values[0]*base^0 + values[1]*base^1 + ... + values[i]*base^i
}

func NewRollingHash(s string) *RollingHash {
	values := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		values[i] = int(s[i])
	}

	inv, cumulativeHashes := constructInvAndHashes(base, values)

	return &RollingHash{
		Base:             base,
		Values:           values,
		Inv:              inv,
		CumulativeHashes: cumulativeHashes,
	}
}

func (rh *RollingHash) Less(l1, r1, l2, r2 int) bool {
	if rh.Find(l1, r1) == rh.Find(l2, r2) {
		return false
	}
	if rh.Values[l1] != rh.Values[l2] {
		return rh.Values[l1] < rh.Values[l2]
	}

	len1 := r1 - l1
	len2 := r2 - l2

	samePrefixCount := 1
	ng := min.Ordered([]int{len1, len2}) + 1

	isPrefixSame := func(k int) bool {
		return rh.Find(l1, l1+k) == rh.Find(l2, l2+k)
	}

	for ng-samePrefixCount > 1 {
		mid := (samePrefixCount + ng) / 2
		if isPrefixSame(mid) {
			samePrefixCount = mid
		} else {
			ng = mid
		}
	}

	if len1 != len2 && samePrefixCount == len1 {
		return true // { 0 <= i < min(|S|, |T|) | Si == Ti } && |S| < |T|
	}
	if len1 != len2 && samePrefixCount == len2 {
		return false // { 0 <= i < min(|S|, |T|) | Si == Ti } && |S| > |T|
	}
	if len1 == len2 && samePrefixCount == len1 {
		return false // { 0 <= i < min(|S|, |T|) | Si == Ti } && |S| == |T|
	}

	return rh.Values[l1+samePrefixCount+1] < rh.Values[l2+samePrefixCount+1]
}

// [l, r)
func (rh *RollingHash) Find(l, r int) int {
	res := rh.CumulativeHashes[r] - rh.CumulativeHashes[l]

	if res < 0 {
		res += Mod
	}

	// k = r - l.
	// (s1*base^l + s2*base^l+1 + ... + sn*base^l+k) -> (s1*base^0+ s2*base^1 + ... + sn*base^k-1)
	return modMul(res, rh.Inv[l])
}

func constructInvAndHashes(base uint, values []int) ([]int, []int) {
	n := len(values)
	inv := make([]int, n+1)
	cumulativeHashes := make([]int, n+1)

	baseInv := int(math.ModInv(base, Mod))
	inv[n] = modPow(baseInv, n)
	basePowI := 1

	for i := 0; i < n; i++ {
		reverseI := n - i - 1
		inv[reverseI] = modMul(int(base), inv[reverseI+1])

		hash := modMul(values[i], basePowI)
		cumulativeHashes[i+1] = mod(hash + cumulativeHashes[i])

		basePowI = modMul(basePowI, int(base))
	}

	return inv, cumulativeHashes
}

func modPow(a, e int) int {
	aPowE := 1
	for e > 0 {
		if e&1 == 1 {
			aPowE = modMul(aPowE, a)
		}

		a = modMul(a, a)
		e >>= 1
	}
	return aPowE
}

func modMul(a, b int) int {
	a1 := a >> 31
	a2 := a & mask31
	b1 := b >> 31
	b2 := b & mask31

	mid := a2*b1 + a1*b2
	mid1 := mid >> 30
	mid2 := mid & mask30

	return mod(a1*b1*2 + mid1 + (mid2 << 31) + a2*b2)
}

func mod(x int) int {
	x1 := x >> 61
	x2 := x & mask61

	res := x1 + x2
	if res >= Mod {
		res -= Mod
	}
	return res
}
