package gcdby

import (
	"math/big"
)

var (
	zero = big.NewInt(0)
	one  = big.NewInt(1)
	two  = big.NewInt(2)
)

func truncate(f *big.Int, t uint) *big.Int {
	if t == 0 {
		return zero
	}
	twoT := new(big.Int).Lsh(one, t-1) // twoT = 1 << (t - 1)
	a := new(big.Int).Add(f, twoT)     // a = f + twoT
	b := new(big.Int).Mul(two, twoT)
	b.Sub(b, one)         // b = 2 * twoT - 1
	a.And(a, b)           // a = a & b
	return a.Sub(a, twoT) // ((f + twoT) & (2 * twoT - 1)) - twoT
}

func getIterations(d uint) uint {
	if d < 46 {
		return (49*d + 80) / 17
	}
	return (49*d + 57) / 17
}
