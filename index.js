function getIterations (digits) {
  return digits < 46
    ? Math.floor((49 * digits + 80) / 17)
    : Math.floor((49 * digits + 57) / 17)
}

function getBitLength (n) {
  return Math.ceil(Math.log2(n + 1))
}

function truncate (f, t) {
  if (!t) return 0
  const twoT = 1 << (t - 1)
  return ((f + twoT) & (2 * twoT - 1)) - twoT
}

function divSteps (initialN, initialT, initialDelta, initialF, initialG) {
  if (initialT < initialN || initialN < 0) {
    throw new Error('invalid parameters for divSteps')
  }
  let n = initialN
  let t = initialT
  let delta = initialDelta
  let f = truncate(initialF, t)
  let g = truncate(initialG, t)
  let u = 1
  let v = 0
  let q = 0
  let r = 1

  while (n > 0) {
    f = truncate(f, t)
    if (delta > 0 && (g & 1)) {
      [delta, f, g, u, v, q, r] = [-delta, g, -f, q, r, -u, -v]
    }
    let g0 = g & 1
    ;[delta, g, q, r] = [1 + delta, (g + g0 * f) / 2, (q + g0 * u) / 2, (r + g0 * v) / 2]
    n -= 1
    t -= 1
    g = truncate(g, t)
  }
  return [delta, f, g, [u, v, q, r]]
}

function gcd (f, g) {
  if (!(f & 1)) {
    throw new Error('getGcd(f, g) only works if f is odd')
  }
  const d = Math.max(getBitLength(f), getBitLength(g))
  const m = getIterations(d)
  const [delta, fm, gm, P] = divSteps(m, m + d, 1, f, g)
  return Math.abs(fm)
}

module.exports = {
  gcd,
  divSteps,
  truncate,
  getIterations,
  getBitLength
}
