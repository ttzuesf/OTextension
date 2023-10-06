package pfield

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
)

var Prinumber = map[string]string{ // map Bits length to Prime number
	"128": "272569594747777388653931295583358822963",
	"160": "1371653710411453064923650500888120813710637140787",
	"256": "95058974245277059354928603316804690503937399282811564918786913047079557400267",
	"384": "34666891571466416369943876924910218769467850733377910570730343977904586843095453473378334307143259769257458062520919",
	"512": "13287065567026334171479949288511240572764681181186512239055368174694960248400076929503278747625210562479783021672131635969335480426986245016254367833939343",
}

type Pfield struct {
	Pr *big.Int `json:"p"` //
	G  *big.Int `json:"g"` // generator element
	E  *big.Int `json:"e"` // unit element
	k  int      `json:"secparam"`
}

func (pf *Pfield) Set(src *big.Int) *big.Int { // assign value
	res := new(big.Int)
	return res.Mod(src, pf.Pr)
}

// a+b mod p
func (pf *Pfield) Add(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Add(a, b)
	return res.Mod(res, pf.Pr)
}

// a-b mod p
func (pf *Pfield) Sub(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Sub(a, b)
	return res.Mod(res, pf.Pr)
}

// a*b mod p
func (pf *Pfield) Mul(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Mul(a, b)
	return res.Mod(res, pf.Pr)
}
func (pf *Pfield) Pow(a *big.Int, num *big.Int) *big.Int {
	res := new(big.Int)
	return res.Exp(a, num, pf.Pr)
}

// a/b mod p
func (pf *Pfield) Div(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int)
	inv := pf.Inverse(b)
	res.Mul(a, inv)
	return res.Mod(res, pf.Pr)
}

// z^{-1} mode p
func (pf *Pfield) Inverse(z *big.Int) *big.Int {
	a := new(big.Int)
	a.Set(z)
	b := new(big.Int)
	b.Set(pf.Pr)
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
		t0.Add(pf.Pr, t0)
	}
	return t0
}
func (pf *Pfield) Double_pow_mul(g1 *big.Int, x *big.Int, g2, y *big.Int) *big.Int { //g1^{x}g2^{y}
	a := pf.Pow(g1, x)
	b := pf.Pow(g2, y)
	return pf.Mul(a, b)
}

func (pf *Pfield) Module() *big.Int {
	res := new(big.Int)
	res.Set(pf.Pr)
	return res
}

func (pf *Pfield) Setbytes(buf []byte) *big.Int {
	res := new(big.Int)
	res.SetBytes(buf)
	return res.Mod(res, pf.Pr)
}

func (pf *Pfield) Sample_field_from_bytes(bitlen int) *big.Int {
	var buf []byte
	if bitlen <= pf.k {
		buf = make([]byte, bitlen)
	} else {
		buf = make([]byte, pf.k)
	}
	rand.Read(buf)
	res := new(big.Int)
	res.SetBytes(buf)
	return res.Mod(res, pf.Pr)
}

// Print() prints field elements
func (pf *Pfield) Print() {
	field, _ := json.Marshal(pf)
	fmt.Printf("The finite field is: %s\n", string(field))
}

// Cmp(a,b) implement the function to compare number a and b
// -1 : a<b
// 0 : a==b
// 1: a>b
func (pf *Pfield) Cmp(a *big.Int, b *big.Int) int { // compare a and b
	return a.Cmp(b)
}
func (pf *Pfield) Init(secparam string) error {
	if _, ok := Prinumber[secparam]; ok != true {
		return errors.New("wrong secure parameter")
	}
	pr := new(big.Int)
	err := json.Unmarshal([]byte(Prinumber[secparam]), pr)
	if err != nil {
		return err
	}
	pf.Pr = pr
	pf.generator()
	pf.E = big.NewInt(1)
	return nil
}

// Generator outputs a generator of Z_{p}
func (pf *Pfield) Generator() *big.Int { // generate the generator
	return pf.G
}

func (pf *Pfield) generator() {
	two := big.NewInt(2)
	Q := new(big.Int)
	Q.Sub(pf.Pr, big.NewInt(1)) // pr-1
	Q.Div(pf.Pr, two)           //(pr-1)/2
	v1 := new(big.Int)
	v2 := new(big.Int)
	for {
		g, _ := rand.Int(rand.Reader, pf.Pr)
		v1.Exp(g, two, pf.Pr) // g^2
		v2.Exp(g, Q, pf.Pr)   // g
		if v1.Cmp(big.NewInt(1)) != 0 && v2.Cmp(big.NewInt(1)) != 1 {
			pf.G = g
			break
		}
	}
}

func NewPfield(secparam string) (*Pfield, error) {
	pf := new(Pfield)
	err := pf.Init(secparam)
	if err != nil {
		return nil, err
	}
	return pf, nil
}

func SavePField(filename string, field *Pfield) error {
	buf, err := json.Marshal(field)
	if err != nil {
		return errors.New("convert field to json wrong")
	}
	log.Println("buf", string(buf))
	err = os.WriteFile(filename, buf, 0666)
	if err != nil {
		return err
	}
	return nil
}

func ImportField(filename string, field *Pfield) error {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return errors.New("can't open specific file")
	}
	err = json.Unmarshal(buf, field)
	if err != nil {
		return errors.New("solve bytes to struct pfield wrong")
	}
	field.k = field.Pr.BitLen()
	return nil
}
