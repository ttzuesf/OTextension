// RsimpleOT is derived the paper named "The Simplest Protocol for Oblivious Transfer",
// Related url: https://eprint.iacr.org/2015/267.pdf
package SimpleOT

import (
	"crypto/rand"
	"errors"
	"github.com/ttzuesf/goot/crypt"
	"github.com/ttzuesf/goot/crypt/AES"
	"github.com/ttzuesf/goot/field/pfield"
	"log"
	"math/big"
	"strconv"
)

type RsimpleOT struct {
	Field *pfield.Pfield
	N     int // message number
}

func (ot *RsimpleOT) Sender(sch, rch chan interface{}, nOTs int, plaintext [][]byte) error {
	pr := ot.Field.Pr
	g := ot.Field.G
	if len(plaintext) != ot.N {
		return errors.New("error plaintext length")
	}
	if len(plaintext) != ot.N {
		return errors.New("error messages")
	}
	for i := 0; i < nOTs; i++ {
		//step 1 A=g^{a}
		a, _ := rand.Int(rand.Reader, pr)
		A := ot.Field.Pow(g, a)
		T := ot.Field.Pow(A, a) // A^{a}
		sch <- A
		// step 2
		msg := <-rch
		B := msg.(*big.Int)
		AB := ot.Field.Pow(B, a)
		ciph := make([][]byte, ot.N)
		for j := 0; j < ot.N; j++ {
			t := big.NewInt(int64(-j))
			//t = ot.Field.Set(t)
			//log.Println("-j", t)
			Tt := ot.Field.Pow(T, t) //T^{-t}=A^{-at}
			r := ot.Field.Mul(Tt, AB)
			k := crypt.Hash(r)
			ciph[j], _ = AES.CBCEncrypt(plaintext[j], k)
		}
		sch <- ciph
		msg = <-rch
	}
	log.Println("Alice: Game Over!")
	return nil
}

func (ot *RsimpleOT) Receiver(sch, rch chan interface{}, nOTs int, choice []int64) {
	pr := ot.Field.Pr
	g := ot.Field.G
	if len(choice) != nOTs {
		return
	}
	c := choice
	ciph := make([][]byte, 0)
	//plain := make([]string, 2)
	for i := 0; i < nOTs; i++ {
		//wait receive A
		msg := <-sch
		A := msg.(*big.Int)
		b, _ := rand.Int(rand.Reader, pr)
		r := big.NewInt(c[i])
		log.Println(r)
		B := ot.Field.Pow(g, b)
		Ar := ot.Field.Pow(A, r)
		R := ot.Field.Mul(Ar, B) //A^{r}g^{b}
		rch <- R
		// Decrypt ciphertext
		ciph0 := <-sch
		ciph = ciph0.([][]byte)
		AB := ot.Field.Pow(A, b)
		k := crypt.Hash(AB)
		plain, err := AES.CBCDecrypt(ciph[c[i]], k)
		log.Println(err)
		log.Println("Extract:", string(plain), "in ", c[i], "th")
		rch <- "End " + strconv.Itoa(i) + "th OT!"
	}
	log.Println("Bob: Game Over")
}

// NewSimpleOT implement an init function to generate new Simple OT
func NewRsimpleOT(field *pfield.Pfield, n int) (*RsimpleOT, error) {
	if field == nil {
		return nil, errors.New("error finite field")
	}
	return &RsimpleOT{Field: field, N: n}, nil
}
