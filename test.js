const assert = require('assert')
const {
  getBitLength,
  getIterations,
  truncate,
  divSteps,
  gcd
} = require('.')

// get iterations
assert.strictEqual(getIterations(10), 33)
assert.strictEqual(getIterations(1), 7)

// getBitLength
assert.strictEqual(getBitLength(100), 7)
assert.strictEqual(getBitLength(10), 4)
assert.strictEqual(getBitLength(1), 1)

// truncate
assert.strictEqual(truncate(100, 5), 4)
assert.strictEqual(truncate(12345, 3), 1)
assert.strictEqual(truncate(1, 1), -1)

// divSteps
assert.deepStrictEqual(divSteps(1, 2, 3, 4, 5), [-2, 1, 0, [0, 1, -1 / 2, 0]])
assert.deepStrictEqual(divSteps(4, 5, 3, 3, 6), [-1, -1, 0, [0, 0.5, -1 / 8, 1 / 16]])
assert.deepStrictEqual(divSteps(36, 47, 1, 903, 1542), [19, -3, 0, [73 / 32768, -213 / 65536, 257 / 34359738368, -301 / 68719476736]])

console.log('Tests pass')
