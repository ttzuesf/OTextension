package SimpleOT

import (
	"github.com/ttzuesf/goot/field/pfield"
	"github.com/ttzuesf/goot/utils"
	"log"
	"strconv"
	"testing"
)

func TestRsimpleOT(t *testing.T) {
	pf := new(pfield.Pfield)
	n := 13
	nOTs := 4 // OT execuation numbers
	pfield.ImportField("field.json", pf)
	ot, err := NewRsimpleOT(pf, n)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("OT", "N=", ot.N)
	plain := make([][]byte, 0)
	for i := 0; i < n; i++ {
		str := "Hello" + strconv.Itoa(i)
		plain = append(plain, []byte(str))
	}
	choice := utils.RandIntvector(nOTs, n) // choice of receiver
	log.Println("choice:", choice)
	sch := make(chan interface{})
	rch := make(chan interface{})
	go func() {
		ot.Sender(sch, rch, nOTs, plain)
	}()
	go func() {
		ot.Receiver(sch, rch, nOTs, choice)
	}()
	select {}
}
