// SimpleOT is derived the paper named "The Simplest Protocol for Oblivious Transfer",
// Related url: https://eprint.iacr.org/2015/267.pdf
package SimpleOT

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"github.com/ttzuesf/goot/crypt/AES"
	"github.com/ttzuesf/goot/field/pfield"
	"github.com/ttzuesf/goot/utils"
	"log"
	"math/big"
	"strconv"
)

type SimpleOT struct {
	Field *pfield.Pfield
}

func (ot *SimpleOT) Sender(sch, rch chan interface{}, nOTs int) {
	pr := ot.Field.Pr
	g := ot.Field.G
	plain := make([][]byte, 2)
	plain[0] = []byte("Hello0")
	plain[1] = []byte("Hello1")
	k := make([][]byte, 2)
	for i := 0; i < nOTs; i++ {
		//step 1 A=g^{a}
		a, _ := rand.Int(rand.Reader, pr)
		A := ot.Field.Pow(g, a)
		sch <- A
		// step 2
		msg := <-rch
		B := msg.(*big.Int)
		// count k_{0}
		AB := ot.Field.Pow(B, a)
		bf, _ := json.Marshal(AB)
		h := sha256.New()
		h.Write(bf)
		k[0] = h.Sum(nil)
		// count k_{1}
		AB1 := ot.Field.Div(B, A)
		AB1 = ot.Field.Pow(AB1, a) // (B/A)^{a}
		bf, _ = json.Marshal(AB1)
		h = sha256.New()
		h.Write(bf)
		k[1] = h.Sum(nil)
		// enc M
		ciph := make([][]byte, 2)
		for j := 0; j < 2; j++ {
			ciph[j], _ = AES.CBCEncrypt(plain[j], k[j])
		}
		sch <- ciph[0]
		sch <- ciph[1]
		msg = <-rch
	}
	log.Println("Alice: Game Over!")
}

func (ot *SimpleOT) Receiver(sch, rch chan interface{}, nOTs int) {
	pr := ot.Field.Pr
	g := ot.Field.G
	//e := ot.Field.E
	var rnd = make([]byte, 8)
	rand.Read(rnd[:7])
	c := utils.Randbitvector(nOTs)
	ciph := make([][]byte, 2)
	//plain := make([]string, 2)
	for i := 0; i < nOTs; i++ {
		//wait receive A
		msg := <-sch
		A := msg.(*big.Int)
		b, _ := rand.Int(rand.Reader, pr)
		B := ot.Field.Pow(g, b)
		if c.Get(i) == true {
			B = ot.Field.Mul(B, A)
		}
		rch <- B
		// Decrypt ciphertext
		ciph0 := <-sch
		ciph[0] = ciph0.([]byte)
		ciph1 := <-sch
		ciph[1] = ciph1.([]byte)
		Ab := ot.Field.Pow(A, b)
		h := sha256.New()
		buf, _ := json.Marshal(Ab)
		h.Write(buf)
		k := h.Sum(nil)
		plaintext := make([]byte, 0) // the final message
		if c.Get(i) == true {
			plaintext, _ = AES.CBCDecrypt(ciph[1], k)
		} else {
			plaintext, _ = AES.CBCDecrypt(ciph[0], k)
		}
		log.Println("Extract:", string(plaintext), "in ", i, "th")
		rch <- "End " + strconv.Itoa(i) + "th OT!"
	}
	log.Println("Bob: Game Over")
}

// NewSimpleOT implement a init function to generate new Simple OT
func NewSimpleOT(field *pfield.Pfield) (*SimpleOT, error) {
	if field == nil {
		return nil, errors.New("error finite field")
	}
	return &SimpleOT{Field: field}, nil
}
