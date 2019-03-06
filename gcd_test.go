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

func TestGetIterations(t *testing.T) {
	type Test struct {
		d        uint
		expected uint
	}
	testCases := []Test{
		Test{d: 10, expected: 33},
		Test{d: 1, expected: 7},
	}
	for _, test := range testCases {
		result := getIterations(test.d)
		if result != test.expected {
			t.Errorf("getIterations(%d) wanted %d, got %d",
				test.d,
				test.expected,
				result,
			)
		}
	}
}
