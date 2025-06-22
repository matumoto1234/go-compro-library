package math_test

import (
	"testing"

	"github.com/matumoto1234/go-compro-library/math"
)

func TestExtGCD(t *testing.T) {
	tests := []struct {
		name    string
		a       int
		b       int
		wantGCD int
		wantX   int
		wantY   int
	}{
		{"111*3 + 30*-11 = 3", 111, 30, 3, 3, -11},
		{"0と数の組", 0, 5, 5, 0, 1},
		{"数と0の組", 7, 0, 7, 1, 0},
		{"互いに素", 35, 64, 1, 11, -6},
		{"負の数含む", -25, 15, 5, 1, 2},   // -25*1 + 15*2 = 5
		{"同じ数", 42, 42, 42, 0, 1},     // 多くの解がある
		{"一方が1", 1, 999, 1, 1, 0},     // x=1, y=0 satisfies 1*1 + 999*0 = 1
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gcd, x, y := math.ExtGCD(tt.a, tt.b)
			if tt.wantGCD != gcd || tt.wantX != x || tt.wantY != y {
				t.Errorf("wantGCD: %d, wantX: %d, wantY: %d, gcd: %d, x: %d, y: %d", tt.wantGCD, tt.wantX, tt.wantY, gcd, x, y)
			}
		})
	}
}
