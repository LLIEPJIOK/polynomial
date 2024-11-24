package polynomial

import (
	"fmt"
	"strconv"
)

func reverse(str string) string {
	rns := []rune(str)

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

type Polynomial struct {
	val int
}

func New(val int) *Polynomial {
	return &Polynomial{
		val: val,
	}
}

func FromStr(str string) (*Polynomial, error) {
	val, err := strconv.ParseInt(reverse(str), 2, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse str to int: %w", err)
	}

	return &Polynomial{
		val: int(val),
	}, nil
}

func (p *Polynomial) String() string {
	str := fmt.Sprintf("%b", p.val)
	return reverse(str)
}

func (p *Polynomial) Deg() int {
	deg := 0
	for ; 1<<(deg+1) <= p.val; deg++ {
	}

	return deg
}

func (p *Polynomial) ToMod(mod *Polynomial) {
	aDeg := p.Deg()
	mDeg := mod.Deg()

	for i := aDeg; i >= mDeg; i-- {
		if p.val&(1<<i) != 0 {
			p.val ^= (mod.val << (i - mDeg))
		}
	}
}

func Add(first, second, mod *Polynomial) *Polynomial {
	ans := New(first.val ^ second.val)
	return ans
}

func Multiply(first, second, mod *Polynomial) *Polynomial {
	ans := New(0)

	for i := 0; 1<<i <= second.val; i++ {
		if second.val&(1<<i) != 0 {
			ans.val ^= first.val << i
		}
	}

	ans.ToMod(mod)
	return ans
}

func Del(first, second, mod *Polynomial) *Polynomial {
	return Multiply(first, Inv(second, mod), mod)
}

func defaultDel(first, second *Polynomial) (*Polynomial, *Polynomial) {
	fDeg, sDeg := first.Deg(), second.Deg()
	ans, rem := New(0), New(first.val)

	for i := fDeg; i >= sDeg; i-- {
		if rem.val&(1<<i) != 0 {
			rem.val ^= (second.val << (i - sDeg))
			ans.val |= (1 << (i - sDeg))
		}
	}

	return ans, rem
}

func Pow(pol *Polynomial, deg int, mod *Polynomial) *Polynomial {
	if deg == 0 {
		return New(1)
	}

	ans := Pow(Multiply(pol, pol, mod), deg/2, mod)
	if deg&1 == 1 {
		ans = Multiply(ans, pol, mod)
	}

	return ans
}

func Inv(pol, mod *Polynomial) *Polynomial {
	dimF := 1 << mod.Deg()
	return Pow(pol, dimF-2, mod)
}

func (p *Polynomial) Reduce() []*Polynomial {
	pCopy := New(p.val)
	pols := make([]*Polynomial, 0)

	for i := 2; i <= pCopy.val; {
		curDel := New(i)
		quotient, rem := defaultDel(pCopy, curDel)

		if rem.val == 0 {
			pols = append(pols, curDel)
			pCopy = quotient
		} else {
			i++
		}
	}

	return pols
}
