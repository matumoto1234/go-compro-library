// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/lesson/1/ALDS1/14/ALDS1_14_D

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"sync"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)
	defer stdout.Flush()

	var t string
	fmt.Fscan(stdin, &t)

	sa := NewSuffixArray(t)

	var q int
	fmt.Fscan(stdin, &q)

	for qi := 0; qi < q; qi++ {
		var p string
		fmt.Fscan(stdin, &p)

		ok := len(t) - 1 // sa[i] >= p なる最小のi
		ng := -1

		isOK := func(x int) bool {
			return t[sa[x]:] >= p
		}

		for ok-ng > 1 {
			mid := (ok + ng) / 2
			if isOK(mid) {
				ok = mid
			} else {
				ng = mid
			}
		}

		if strings.HasPrefix(t[sa[ok]:], p) {
			fmt.Fprintln(stdout, 1)
		} else {
			fmt.Fprintln(stdout, 0)
		}
	}
}

func NewSuffixArray(s string) []int {
	rh := NewRollingHash(s)

	sa := make([]int, len(s))
	for i := range sa {
		sa[i] = i
	}

	sort.Slice(sa, func(i, j int) bool {
		return rh.Less(sa[i], len(s), sa[j], len(s))
	})

	return sa
}

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

type Float interface {
	~float32 | ~float64
}

type Ordered interface {
	Integer | Float | ~string
}

func min[T Ordered](a []T) T {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
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
	ng := min([]int{len1, len2}) + 1

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

	return rh.Values[l1+samePrefixCount] < rh.Values[l2+samePrefixCount]
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

	baseInv := int(ModInv(base, Mod))
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

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// return (gcd, x, y)
func ExtGCD[T Integer](a, b T) (T, T, T) {
	return extGCD(a, b, 0, 0)
}

func extGCD[T Integer](a, b, x, y T) (T, T, T) {
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

// a^-1 (mod m)
func ModInv[T Integer](a, m T) T {
	// a*x + mod*y = 1
	_, x, _ := ExtGCD(a, m)
	if x < 0 {
		x += m
	}
	return x % m
}
