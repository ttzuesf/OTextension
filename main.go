package main

import (
	"crypto/rand"
	"encoding/binary"
	"log"
)

func main() {
	a := make([]byte, 8)
	rand.Read(a)
	log.Println(a)
	b := binary.LittleEndian.Uint64(a)
	//c := (*int64)(unsafe.Pointer(&b))
	log.Println(b)
}
