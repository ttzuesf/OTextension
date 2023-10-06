package AES

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"log"
	"testing"
)

func TestCBCAES(t *testing.T) {
	m := []byte("hello")
	//rand.Read(m)
	k := make([]byte, 32)
	rand.Read(k)
	c, err := CBCEncrypt(m, k)
	if err != nil {
		log.Fatal(err)
	}
	msg, err := CBCDecrypt(c, k)
	if bytes.Compare(msg, m) == 0 {
		log.Println("m:", string(m))
		log.Println("msg:", string(msg))
		fmt.Println("decrypt ciphertext error")
	}
}
