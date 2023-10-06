package integer

import (
	"log"
	"testing"
)

func TestGeneratePrimeNumber(t *testing.T) {
	prim := GeneratePrimeNumber(224)
	pr := EncodeNumber(prim)
	log.Println("Encoded prime:", pr)
	prim1, err := DecodeNumber(pr)
	if err != nil {
		log.Fatal("Decode wrong!")
	}
	if prim1.Cmp(prim) == 0 {
		log.Println("Decode correctly")
	}
}
