package gcdby

import (
	"math/big"
)

type step struct {
	delta int
	f     *big.Int
	g     *big.Int
	p     []*big.Float
}

var (
	zero  = big.NewInt(0)
	zerof = big.NewFloat(0)
	one   = big.NewInt(1)
	onef  = big.NewFloat(1)
	two   = big.NewInt(2)
	twof  = big.NewFloat(2)
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

func iterations(d uint) uint {
	if d < 46 {
		return (49*d + 80) / 17
	}
	return (49*d + 57) / 17
}

func divstep(n, t uint, delta int, initialF, initialG *big.Int) step {
	if t < n || n < 0 {
		panic("invalid divstep arguments")
	}
	f, g := truncate(initialF, t), truncate(initialG, t)
	u, v, q, r := big.NewFloat(1), big.NewFloat(0), big.NewFloat(0), big.NewFloat(1)

	for n > 0 {
		f = truncate(f, t)
		g0 := new(big.Int).And(g, one)
		if delta > 0 && (g0.Cmp(one) == 0) {
			delta, f, g, u, v, q, r = -delta, g, new(big.Int).Neg(f), q, r, new(big.Float).Neg(u), new(big.Float).Neg(v)
		}
		g0.And(g, one)
		g0f := new(big.Float).SetInt(g0)
		delta = 1 + delta
		g.Add(g, new(big.Int).Mul(g0, f)).Quo(g, two)     // (g+g0*f)/2
		q.Add(q, new(big.Float).Mul(g0f, u)).Quo(q, twof) // (q+g0f*u)/2
		r.Add(r, new(big.Float).Mul(g0f, v)).Quo(r, twof) // (r+g0f*v)/2
		n, t = n-1, t-1
		g = truncate(g, t)
	}

	result := step{
		delta: delta,
		f:     f,
		g:     g,
		p:     []*big.Float{u, v, q, r},
	}
	return result
}

func GcdInt(f, g int) int {
	return 0
}

func Gcd(f, g *big.Int) {
	f1 := new(big.Int).And(f, one)
	if f1.Cmp(one) == 0 {
		panic("f must be odd")
	}

	// const d = Math.max(getBitLength(f), getBitLength(g))
	// const m = getIterations(d)
	// const [delta, fm, gm, P] = divSteps(m, m + d, 1, f, g)
	// return Math.abs(fm)
}
