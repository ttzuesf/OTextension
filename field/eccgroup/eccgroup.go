package eccgroup

import (
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/ttzuef/goot/field/Zn"
	"math/big"
	"strconv"
)

// Point

type Point struct {
	X *big.Int
	Y *big.Int
}

type Curve struct {
	curv elliptic.Curve
	P    *big.Int // A prime number to construct Fp
	N    *big.Int // The order of elliptic curve group
	zero *big.Int
	G    *Point //  The base point
	Zp   *Zn.Zn `json:"zp"`
}

// Assign a point to another varible
func (ecf *Curve) Set(P *Point) *Point { // assign value
	if P.X == nil || P.Y == nil {
		return nil
	}
	if !ecf.curv.IsOnCurve(P.X, P.Y) {
		return nil
	}
	return &Point{
		X: P.X,
		Y: P.Y,
	}
}

// P+Q
func (ecf *Curve) Mul(P, Q *Point) *Point {
	res := new(Point)
	res.X, res.Y = ecf.curv.Add(P.X, P.Y, Q.X, Q.Y)
	return res
}

// P-Q
func (ecf *Curve) Div(P, Q *Point) *Point {
	res := new(Point)
	Invq := ecf.Inverse(Q)
	res.X, res.Y = ecf.curv.Add(P.X, P.Y, Invq.X, Invq.Y)
	return res
}

// nP
func (ecf *Curve) Pow(P *Point, n *big.Int) *Point {
	p := new(Point)
	if n.Cmp(ecf.zero) < 0 {
		n.Mod(n, ecf.P)
	}
	p.X, p.Y = ecf.curv.ScalarMult(P.X, P.Y, n.Bytes())
	return p
}

// n*G
func (ecf *Curve) PowG(n *big.Int) *Point {
	p := new(Point)
	if n.Cmp(ecf.zero) < 0 {
		n.Mod(n, ecf.P)
	}
	p.X, p.Y = ecf.curv.ScalarBaseMult(n.Bytes())
	return p
}

// -(x,y)
func (ecf *Curve) Inverse(P *Point) *Point {
	x := new(big.Int).Set(P.X)
	y := new(big.Int).Set(P.Y)
	y.Sub(ecf.zero, y)
	y.Mod(y, ecf.P)
	return &Point{
		X: x,
		Y: y,
	}
}

// 2P
func (ecf *Curve) Double(x, y *big.Int) (*big.Int, *big.Int) {
	return ecf.curv.Double(x, y)
}

// sample a number
func (ecf *Curve) SamplePoint() (*big.Int, *Point) {
	res, _ := rand.Int(rand.Reader, ecf.N)
	return res, ecf.PowG(res)
}

// Print() prints field elements
func (ecf *Curve) Print() {
}

// Cmp(a,b) implement the function to compare number a and b
// 1 : a!=b
// 0 : a==b
func (ecf *Curve) Cmp(P, Q *Point) int { // compare a and b
	a := P.X.Cmp(Q.X)
	b := P.Y.Cmp(Q.Y)
	if a != 0 || b != 0 {
		return 1
	}
	return 0
}
func (ecf *Curve) init(ctype string) error {
	switch ctype {
	case "P224":
		ecf.curv = elliptic.P224()
	case "P256":
		ecf.curv = elliptic.P256()
	case "P384":
		ecf.curv = elliptic.P384()
	case "P521":
		ecf.curv = elliptic.P521()
	default:
		return errors.New("error curve type")
	}
	param := ecf.curv.Params()
	ecf.P = param.P
	ecf.N = param.N
	ecf.G = new(Point)
	ecf.G.X = new(big.Int).Set(param.Gx)
	ecf.G.Y = new(big.Int).Set(param.Gy)
	ecf.zero = big.NewInt(0)
	return nil
}

// Generator outputs the base Point of ECC
func (ecf *Curve) Generator() *Point { // return the generator
	res := new(Point)
	res.X, res.Y = ecf.G.X, ecf.G.Y
	return res
}

// Module returns the module number
func (ecf *Curve) Module() *big.Int { // return the generator
	res := new(big.Int)
	res.Set(ecf.P)
	return res
}

// Order returns the order of elliptic curve group

func (ecf *Curve) Order() *big.Int { // return the generator
	res := new(big.Int)
	res.Set(ecf.N)
	return res
}

func NewECC(secparam int) *Curve {
	res := new(Curve)
	res.init("P" + strconv.Itoa(secparam))
	return res
}
