package bitvector

import (
	"crypto/rand"
	"encoding/binary"
	"math/bits"
)

const bitsPerUnit = bits.UintSize

type Bitsvect []uint

type BitOperation struct {
}

// x xor y
func Xor(x, y Bitsvect) Bitsvect {
	const e = 1
	m := len(x)
	n := len(y)
	if m > n {
		m, n = n, m
	}
	res := make(Bitsvect, n, n+e)
	for i := 0; i < m; i++ {
		res[i] = x[i] ^ y[i]
	}
	if len(x) >= len(y) {
		copy(res[m:n], x[m:n])
	} else {
		copy(res[m:n], y[m:n])
	}
	return res
}

// x bitwise AND y
func And(x, y Bitsvect) Bitsvect {
	const e = 1
	m := len(x)
	n := len(y)
	if m > n {
		m, n = n, m
	}
	res := make(Bitsvect, n, n+e)
	for i := 0; i < m; i++ {
		res[i] = x[i] & y[i]
	}
	return res
}

// x or y

func Or(x, y Bitsvect) Bitsvect {
	const e = 1
	m := len(x)
	n := len(y)
	if m > n {
		m, n = n, m
	}
	res := make(Bitsvect, n, n+e)
	for i := 0; i < m; i++ {
		res[i] = x[i] | y[i]
	}
	if len(x) >= len(y) {
		copy(res[m:n], x[m:n])
	} else {
		copy(res[m:n], y[m:n])
	}
	return res
}

// Set a bits vector
func Set(x []bool) Bitsvect {
	const e = 4
	n := len(x)
	l := n / bitsPerUnit
	r := n % bitsPerUnit
	m := l
	if r != 0 {
		m = l + 1
	}
	res := make(Bitsvect, m, n+e)
	for k, v := range x {
		if v {
			res[k/bitsPerUnit] |= (1 << (k % bitsPerUnit))
		} else {
			res[k/bitsPerUnit] &= ^(1 << (k % bitsPerUnit))
		}
	}
	return res
}

//

func AtoBits(x string) Bitsvect {
	const e = 4
	y := []byte(x)
	n := len(y)
	l := n / bitsPerUnit
	r := n % bitsPerUnit
	m := l
	if r != 0 {
		m = l + 1
	}
	res := make(Bitsvect, m, n+e)
	for k, v := range y {
		if v == '1' {
			res[k/bitsPerUnit] |= (1 << (k % bitsPerUnit))
		} else {
			res[k/bitsPerUnit] &= ^(1 << (k % bitsPerUnit))
		}
	}
	return res
}

// Set a bits vector

func Bools(x Bitsvect) []bool {
	n := len(x)
	m := (n - 1) * bitsPerUnit
	res := make([]bool, m, m+bitsPerUnit)
	for i := 0; i < n-1; i++ {
		y := x[i]
		j := 0
		for y != 0 {
			k := y & 1
			if k == 0 {
				res[i*bitsPerUnit+j] = false
			} else {
				res[i*bitsPerUnit+j] = true
			}
			j++
			y >>= 1
		}
	}
	y := x[n-1]
	for y != 0 {
		k := y & 1
		if k == 0 {
			res = append(res, false)
		} else {
			res = append(res, true)
		}
		y >>= 1
	}
	return res
}

// BitstoA convert bits vect to a bits string
func BitstoA(x Bitsvect) string {
	n := len(x)
	m := (n - 1) * bitsPerUnit
	res := make([]byte, m, m+bitsPerUnit)
	for i := 0; i < n-1; i++ {
		y := x[i]
		j := 0
		for y != 0 {
			k := y & 1
			if k == 0 {
				res[i*bitsPerUnit+j] = byte('0')
			} else {
				res[i*bitsPerUnit+j] = byte('1')
			}
			j++
			y >>= 1
		}
	}
	y := x[n-1]
	for y != 0 {
		k := y & 1
		if k == 0 {
			res = append(res, byte('0'))
		} else {
			res = append(res, byte('1'))
		}
		y >>= 1
	}
	return string(res)
}

// Extract a bit from a bits vector

func Extract(a Bitsvect, index int) bool {
	n := index / bitsPerUnit
	ind := index - n*bitsPerUnit
	val := a[n] & (1 << ind)
	if val == 0 {
		return false
	}
	return true
}

func Length(x Bitsvect) int {
	n := len(x)
	y := x[n-1]
	s := 0
	for y != 0 {
		y = y >> 1
		s++
	}
	return (n-1)*bitsPerUnit + s
}

// RandBitsvect implements a algorithm to produce a random bits vector

func RandBitsvect(bitlen int) Bitsvect {
	l := bitlen / bitsPerUnit
	r := bitlen % bitsPerUnit
	if l*bitsPerUnit < bitlen {
		l = l + 1
	}
	buf := make([]byte, l*8)
	rand.Read(buf)
	var res Bitsvect
	for i := 0; i < l; i++ {
		a := binary.BigEndian.Uint64(buf[i*8 : 8*(i+1)])
		res = append(res, uint(a))
	}
	if r != 0 {
		res[len(res)-1] >>= bitsPerUnit - r
	}
	return res
}
