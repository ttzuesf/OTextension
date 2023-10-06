package SimpleOT

import (
	"crypto/rand"
	"github.com/ttzuesf/goot/field/pfield"
	"log"
	"sync"
	"testing"
	"unsafe"
)

func TestSimpleOT(t *testing.T) {
	pf := new(pfield.Pfield)
	pfield.ImportField("field.json", pf)
	log.Println("Start simplest OT game between Alice and Bob:")
	log.Println("Initial field:", pf)
	simot, err := NewSimpleOT(pf)
	if err != nil {
		log.Fatal(err)
	}
	sch := make(chan interface{})
	rch := make(chan interface{})
	nOTs := 1024
	go simot.Sender(sch, rch, nOTs)
	go simot.Receiver(sch, rch, nOTs)
	select {}
}

func BenchmarkSimpleOT(b *testing.B) {
	pf := new(pfield.Pfield)
	pfield.ImportField("field.json", pf)
	log.Println("Start simplest OT game between Alice and Bob:")
	log.Println("Initial field:", pf)
	simot, err := NewSimpleOT(pf)
	if err != nil {
		log.Fatal(err)
	}
	sch := make(chan interface{})
	rch := make(chan interface{})
	nOTs := 80
	for i := 0; i < b.N; i++ {
		wg := new(sync.WaitGroup)
		wg.Add(2)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			simot.Sender(sch, rch, nOTs)
		}(wg)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			simot.Receiver(sch, rch, nOTs)
		}(wg)
		wg.Wait()
	}
}

func TestUnsafeOT(t *testing.T) {
	a := make([]byte, 8)
	rand.Read(a[:7])
	log.Println(a)
	c := (*int64)(unsafe.Pointer(&a[0]))
	log.Println(*c)
	//c = nil
}
