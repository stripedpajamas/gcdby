package gcdby

import (
	"math/big"
	"testing"
)

func TestTruncate(t *testing.T) {
	type Test struct {
		f        *big.Int
		t        uint
		expected *big.Int
	}
	testCases := []Test{
		Test{f: big.NewInt(369), t: 41, expected: big.NewInt(369)},
		Test{f: big.NewInt(100), t: 5, expected: big.NewInt(4)},
		Test{f: big.NewInt(12345), t: 3, expected: big.NewInt(1)},
	}
	for _, test := range testCases {
		result := truncate(test.f, test.t)
		if result.Cmp(test.expected) != 0 {
			t.Errorf("truncate(%s, %d) wanted %s, got %s",
				test.f.String(),
				test.t,
				test.expected.String(),
				result.String(),
			)
		}
	}
}

func TestIterations(t *testing.T) {
	type Test struct {
		d        uint
		expected uint
	}
	testCases := []Test{
		Test{d: 10, expected: 33},
		Test{d: 1, expected: 7},
	}
	for _, test := range testCases {
		result := iterations(test.d)
		if result != test.expected {
			t.Errorf("getIterations(%d) wanted %d, got %d",
				test.d,
				test.expected,
				result,
			)
		}
	}
}

func TestDivstep(t *testing.T) {
	type Test struct {
		n        uint
		t        uint
		delta    int
		initialF *big.Int
		initialG *big.Int
		expected step
	}
	testCases := []Test{
		Test{
			n: 1, t: 2, delta: 3, initialF: big.NewInt(4), initialG: big.NewInt(5),
			expected: step{delta: -2, f: big.NewInt(1), g: big.NewInt(0), p: []*big.Float{
				zerof, onef, big.NewFloat(-0.5), zerof,
			}},
		},
		Test{
			n: 36, t: 47, delta: 1, initialF: big.NewInt(903), initialG: big.NewInt(1542),
			expected: step{delta: 19, f: big.NewInt(-3), g: big.NewInt(0), p: []*big.Float{
				big.NewFloat(73 / 32768), big.NewFloat(-213 / 65536), big.NewFloat(257 / 34359738368), big.NewFloat(-301 / 68719476736),
			}},
		},
	}
	equalStep := func(a, b step) bool {
		if a.delta != b.delta || a.f.Cmp(b.f) != 0 || a.g.Cmp(b.g) != 0 {
			return false
		}
		for i, fa := range a.p {
			fb := a.p[i]
			if fa.Cmp(fb) != 0 {
				return false
			}
		}
		return true
	}
	for _, test := range testCases {
		result := divstep(test.n, test.t, test.delta, test.initialF, test.initialG)
		if !equalStep(result, test.expected) {
			t.Errorf(`divstep wanted
step.delta: %d; got %d
step.f: %s; got %s
step.g: %s; got %s
step.p: [%s, %s, %s, %s]; got [%s, %s, %s, %s]
				`,
				test.expected.delta,
				result.delta,
				test.expected.f.String(),
				result.f.String(),
				test.expected.g.String(),
				result.g.String(),
				test.expected.p[0].String(),
				test.expected.p[1].String(),
				test.expected.p[2].String(),
				test.expected.p[3].String(),
				result.p[0].String(),
				result.p[1].String(),
				result.p[2].String(),
				result.p[3].String(),
			)
		}
	}
}

func TestGcd(t *testing.T) {
	type Test struct {
		f        *big.Int
		g        *big.Int
		expected *big.Int
	}
	testCases := []Test{
		Test{f: big.NewInt(1542), g: big.NewInt(903), expected: big.NewInt(3)},
	}
	for _, test := range testCases {
		result := Gcd(test.f, test.g)
		if result.Cmp(test.expected) != 0 {
			t.Errorf("Gcd(%s, %s) wanted %s, got %s",
				test.f.String(),
				test.g.String(),
				test.expected.String(),
				result.String(),
			)
		}
	}
}
