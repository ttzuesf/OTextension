// Zn implements an integer ring, the N is prime number Zn is a prime field
package Zn

import (
	"crypto/rand"
	"math/big"
)

var Prinumber = map[string]string{ // map Bits length to Prime number
	"128": "272569594747777388653931295583358822963",
	"160": "1371653710411453064923650500888120813710637140787",
	"256": "95058974245277059354928603316804690503937399282811564918786913047079557400267",
	"384": "34666891571466416369943876924910218769467850733377910570730343977904586843095453473378334307143259769257458062520919",
	"512": "13287065567026334171479949288511240572764681181186512239055368174694960248400076929503278747625210562479783021672131635969335480426986245016254367833939343",
}

type Zn struct {
	N    *big.Int `json:"n"` //
	Bits int      `json:"bits"`
}

func (zn *Zn) Set(src *big.Int) *big.Int { // assign value
	res := new(big.Int)
	return res.Mod(src, zn.N)
}

// a+b mod p
func (zn *Zn) Add(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Add(a, b)
	return res.Mod(res, zn.N)
}

// a-b mod n
func (zn *Zn) Sub(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Sub(a, b)
	return res.Mod(res, zn.N)
}

// a*b mod n
func (zn *Zn) Mul(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Mul(a, b)
	return res.Mod(res, zn.N)
}
func (zn *Zn) Pow(a *big.Int, num *big.Int) *big.Int {
	res := new(big.Int)
	return res.Exp(a, num, zn.N)
}

// a/b mod p
func (z *Zn) Div(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int)
	inv := z.Inverse(b)
	res.Mul(a, inv)
	return res.Mod(res, z.N)
}

// z^{-1} mod N
func (zn *Zn) Inverse(z *big.Int) *big.Int {
	a := new(big.Int)
	a.Set(z)
	b := new(big.Int)
	b.Set(zn.N)
	zero := big.NewInt(0)
	rem := new(big.Int)
	quo := new(big.Int)
	s := new(big.Int)
	t := new(big.Int)
	if a.Cmp(b) == -1 { // a< b
		t.Set(a)
		a.Set(b)
		b.Set(t)
	}
	s0 := big.NewInt(1)
	t0 := big.NewInt(0)
	s1 := big.NewInt(0)
	t1 := big.NewInt(1)
	for b.Cmp(zero) == 1 {
		rem.Mod(a, b) // a mod b
		quo.Div(a, b) // a div b
		s.Mul(quo, s1)
		s.Sub(s0, s) // s = s0-qs1
		t.Mul(quo, t1)
		t.Sub(t0, t) // t=t0-qt1
		s0.Set(s1)
		s1.Set(s)
		t0.Set(t1)
		t1.Set(t)
		a.Set(b)
		b.Set(rem)
	}
	if t0.Cmp(zero) == -1 {
		t0.Add(zn.N, t0)
	}
	return t0
}
func (zn *Zn) Double_pow_mul(g1 *big.Int, x *big.Int, g2, y *big.Int) *big.Int { //g1^{x}g2^{y}
	a := zn.Pow(g1, x)
	b := zn.Pow(g2, y)
	return zn.Mul(a, b)
}

func (zn *Zn) Module() *big.Int {
	res := new(big.Int)
	res.Set(zn.N)
	return res
}

func (zn *Zn) Setbytes(buf []byte) *big.Int {
	res := new(big.Int)
	res.SetBytes(buf)
	return res.Mod(res, zn.N)
}

func (zn *Zn) SampleNumber(bitlen int) *big.Int {
	var buf []byte
	if bitlen <= zn.Bits {
		buf = make([]byte, bitlen)
	} else {
		buf = make([]byte, zn.Bits)
	}
	rand.Read(buf)
	res := new(big.Int)
	res.SetBytes(buf)
	return res.Mod(res, zn.N)
}

// Cmp(a,b) implement the function to compare number a and b
// -1 : a<b
// 0 : a==b
// 1: a>b
func (zn *Zn) Cmp(a *big.Int, b *big.Int) int { // compare a and b
	return a.Cmp(b)
}
