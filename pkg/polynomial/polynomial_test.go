package polynomial_test

import (
	"testing"

	"github.com/LLIEPJIOK/polynomial/pkg/polynomial"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromStr(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name string
		str  string
		val  int
	}

	tt := []TestCase{
		{
			name: "single character",
			str:  "1",
			val:  0b1,
		},
		{
			name: "three characters",
			str:  "011",
			val:  0b110,
		},
		{
			name: "multiple characters",
			str:  "001001110010110010100111",
			val:  0b111001010011010011100100,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p, err := polynomial.FromStr(tc.str)
			require.NoError(t, err)

			assert.Equal(t, polynomial.New(tc.val), p)
		})
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name string
		val  int
		str  string
	}

	tt := []TestCase{
		{
			name: "single character",
			val:  0b1,
			str:  "1",
		},
		{
			name: "three characters",
			val:  0b110,
			str:  "011",
		},
		{
			name: "multiple characters",
			val:  0b111001010011010011100100,
			str:  "001001110010110010100111",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := polynomial.New(tc.val)
			assert.Equal(t, tc.str, p.String())
		})
	}
}

func TestToMod(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name string
		val  int
		mod  int
		exp  int
	}

	tt := []TestCase{
		{
			name: "deg=2",
			val:  0b111,
			mod:  0b11,
			exp:  0b1,
		},
		{
			name: "deg=4",
			val:  0b10110,
			mod:  0b101,
			exp:  0b10,
		},
		{
			name: "deg=6",
			val:  0b1001100,
			mod:  0b1101,
			exp:  0b111,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := polynomial.New(tc.val)
			p.ToMod(polynomial.New(tc.mod))
			assert.Equal(t, polynomial.New(tc.exp), p)
		})
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name   string
		first  int
		second int
		mod    int
		exp    int
	}

	tt := []TestCase{
		{
			name:   "deg<=2",
			first:  0b1,
			second: 0b11,
			mod:    0b10000,
			exp:    0b10,
		},
		{
			name:   "deg<=4",
			first:  0b1010,
			second: 0b0101,
			mod:    0b10000,
			exp:    0b1111,
		},
		{
			name:   "deg<=6",
			first:  0b1111,
			second: 0b101111,
			mod:    0b10000000,
			exp:    0b100000,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			firstPol, secondPol := polynomial.New(tc.first), polynomial.New(tc.second)
			res := polynomial.Add(firstPol, secondPol, polynomial.New(tc.mod))
			assert.Equal(t, polynomial.New(tc.exp), res)
		})
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name   string
		first  int
		second int
		mod    int
		exp    int
	}

	tt := []TestCase{
		{
			name:   "deg<=2",
			first:  0b1,
			second: 0b11,
			mod:    0b111,
			exp:    0b11,
		},
		//    1010
		//  101000
		//  100010
		//  000100
		{
			name:   "deg<=4",
			first:  0b1010,
			second: 0b0101,
			mod:    0b10011,
			exp:    0b100,
		},
		//    101111
		//   101111
		//  101111
		// 101111
		// 110110101
		// 011101001
		//  01000111
		//   0010000
		{
			name:   "deg<=6",
			first:  0b1111,
			second: 0b101111,
			mod:    0b1010111,
			exp:    0b10000,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			firstPol, secondPol := polynomial.New(tc.first), polynomial.New(tc.second)
			res := polynomial.Multiply(firstPol, secondPol, polynomial.New(tc.mod))
			assert.Equal(t, polynomial.New(tc.exp), res)
		})
	}
}

func TestDel(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name   string
		first  int
		second int
		mod    int
		exp    int
	}

	tt := []TestCase{
		{
			name:   "deg<=2",
			first:  0b11,
			second: 0b11,
			mod:    0b111,
			exp:    0b1,
		},
		{
			name:   "deg<=4",
			first:  0b100,
			second: 0b0101,
			mod:    0b10011,
			exp:    0b1010,
		},
		{
			name:   "deg<=6",
			first:  0b10000,
			second: 0b101111,
			mod:    0b1010111,
			exp:    0b1111,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			firstPol, secondPol := polynomial.New(tc.first), polynomial.New(tc.second)
			res := polynomial.Del(firstPol, secondPol, polynomial.New(tc.mod))
			assert.Equal(t, polynomial.New(tc.exp), res)
		})
	}
}

func TestPow(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name   string
		first  int
		second int
		mod    int
		exp    int
	}

	tt := []TestCase{
		{
			name:   "deg<=2",
			first:  0b11,
			second: 2,
			mod:    0b111,
			exp:    0b10,
		},
		{
			name:   "deg<=4",
			first:  0b1010,
			second: 3,
			mod:    0b10101,
			exp:    0b1000,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			firstPol := polynomial.New(tc.first)
			res := polynomial.Pow(firstPol, tc.second, polynomial.New(tc.mod))
			assert.Equal(t, polynomial.New(tc.exp), res)
		})
	}
}

func TestInv(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name string
		val  int
		mod  int
		exp  int
	}

	tt := []TestCase{
		{
			name: "deg<=2",
			val:  0b11,
			mod:  0b111,
			exp:  0b10,
		},
		{
			name: "deg<=4",
			val:  0b1010,
			mod:  0b10011,
			exp:  0b1100,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := polynomial.Inv(polynomial.New(tc.val), polynomial.New(tc.mod))
			assert.Equal(t, polynomial.New(tc.exp), res)
		})
	}
}

func TestReduce(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name string
		val  int
		exp  []*polynomial.Polynomial
	}

	tt := []TestCase{
		{
			name: "deg<=2",
			val:  0b101,
			exp:  []*polynomial.Polynomial{polynomial.New(0b11), polynomial.New(0b11)},
		},
		{
			name: "deg<=6",
			val:  0b110001,
			exp:  []*polynomial.Polynomial{polynomial.New(0b111), polynomial.New(0b1011)},
		},
		{
			name: "deg<=24",
			val:  0b1000000000000000000000001,
			exp: []*polynomial.Polynomial{
				polynomial.New(0b11),
				polynomial.New(0b11),
				polynomial.New(0b11),
				polynomial.New(0b11),
				polynomial.New(0b11),
				polynomial.New(0b11),
				polynomial.New(0b11),
				polynomial.New(0b11),
				polynomial.New(0b111),
				polynomial.New(0b111),
				polynomial.New(0b111),
				polynomial.New(0b111),
				polynomial.New(0b111),
				polynomial.New(0b111),
				polynomial.New(0b111),
				polynomial.New(0b111),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			p := polynomial.New(tc.val)
			assert.ElementsMatch(t, tc.exp, p.Reduce())
		})
	}
}
