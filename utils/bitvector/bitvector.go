package bitvector

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/bits"
)

const bitsPerUnit = bits.UintSize

type Bitsvect struct {
	array []uint
	n     int // bitslength
}

// x xor y
func (v *Bitsvect) Xor(x, y *Bitsvect) *Bitsvect {
	const e = 1
	m := len(x.array)
	n := len(y.array)
	if m > n {
		m, n = n, m
	}
	v.array = make([]uint, n, n+e)
	for i := 0; i < m; i++ {
		v.array[i] = x.array[i] ^ y.array[i]
	}
	if x.n >= y.n {
		copy(v.array[m:n], x.array[m:n])
	} else {
		copy(v.array[m:n], y.array[m:n])
	}
	v.n = y.n
	if x.n > y.n {
		v.n = x.n
	}
	return v
}

// x bitwise AND y
func (v *Bitsvect) And(x, y *Bitsvect) *Bitsvect {
	const e = 1
	m := len(x.array)
	n := len(y.array)
	if m > n {
		m, n = n, m
	}
	v.array = make([]uint, n, n+e)
	for i := 0; i < m; i++ {
		v.array[i] = x.array[i] & y.array[i]
	}
	v.n = y.n
	if x.n > y.n {
		v.n = x.n
	}
	return v
}

// x or y

func (v *Bitsvect) Or(x, y *Bitsvect) *Bitsvect {
	const e = 1
	m := len(x.array)
	n := len(y.array)
	if m > n {
		m, n = n, m
	}
	v.array = make([]uint, n, n+e)
	for i := 0; i < m; i++ {
		v.array[i] = x.array[i] | y.array[i]
	}
	if x.n >= y.n {
		copy(v.array[m:n], v.array[m:n])
	} else {
		copy(v.array[m:n], v.array[m:n])
	}
	v.n = y.n
	if x.n > y.n {
		v.n = x.n
	}
	return v
}

// Set a bits vector
func (v *Bitsvect) Set(str string) {
	if v == nil {
		v = new(Bitsvect)
	}
	v.array, v.n = atobits(str)
}

// Set a bits vector

// BitstoA convert bits vect to a bits string
func (v *Bitsvect) String() string {
	str := make([]byte, v.n)
	j := v.n - 1
	for i := 0; i < v.n; i++ {
		k := i / bitsPerUnit
		l := i - k*bitsPerUnit
		bit := v.array[k] & (1 << l)
		str[j] = '1'
		if bit == 0 {
			str[j] = '0'
		}
		j--
	}
	return string(str)
}

// Extract a bit from a bits vector

func (v *Bitsvect) Extract(a *Bitsvect, index int) byte {
	n := index / bitsPerUnit
	ind := index - n*bitsPerUnit
	val := a.array[n] & (1 << ind)
	return byte(val)
}

func (v *Bitsvect) Length(x *Bitsvect) int {
	return x.n
}

func (v *Bitsvect) Bits(x *Bitsvect) []uint {
	return x.array
}

// RandBitsvect implements a algorithm to produce a random bits vector

func SampleBitsvect(bitlen int) *Bitsvect {
	v := new(Bitsvect)
	l := bitlen / bitsPerUnit
	r := bitlen % bitsPerUnit
	if l*bitsPerUnit < bitlen {
		l = l + 1
	}
	buf := make([]byte, l*8)
	rand.Read(buf)
	var res []uint
	for i := 0; i < l; i++ {
		a := binary.BigEndian.Uint64(buf[i*8 : 8*(i+1)])
		res = append(res, uint(a))
	}
	if r != 0 {
		res[len(res)-1] >>= bitsPerUnit - r
	}
	v.array = res
	v.n = bitlen
	return v
}

// Print

func (v *Bitsvect) Print(a *Bitsvect) {
	str := make([]byte, a.n)
	j := a.n - 1
	for i := 0; i < a.n; i++ {
		k := i / bitsPerUnit
		l := i - k*bitsPerUnit
		bit := a.array[k] & (1 << l)
		str[j] = '1'
		if bit == 0 {
			str[j] = '0'
		}
		j--
	}
	fmt.Printf("%s\n", string(str))
}

// Initial a bits vector from a string

func NewBitsvector(num string) *Bitsvect {
	res := new(Bitsvect)
	res.array, res.n = atobits(num)
	return res
}

//

func atobits(x string) ([]uint, int) {
	const e = 4
	n := len(x)
	n1 := n/bitsPerUnit + 1
	if n%bitsPerUnit == 0 {
		n1 = n / bitsPerUnit
	}
	res := make([]uint, n1, n+e)
	j := 0
	for i := n - 1; i >= 0; i-- {
		k1 := j / bitsPerUnit
		l := j - k1*bitsPerUnit
		if x[i] == '1' {
			res[k1] |= (1 << l)
		}
		j++
	}
	return res, n
}
