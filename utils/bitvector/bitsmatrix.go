package bitvector

// mat=(a_{1},a_{2},a_{3}, \cdots, a_{n})
// a_{i}=m

type Matrix struct {
	array []*Bitsvect
	m     int // rows
	n     int // columns
}

func (mat *Matrix) Xor(a, b *Matrix) *Matrix {
	if a.m != b.m && a.n != b.n {
		return nil
	}
	res := make([]*Bitsvect, b.n)
	for i := 0; i < b.n; i++ {
		c := new(Bitsvect)
		c.Xor(a.array[i], b.array[i])
		res[i] = c
	}
	mat.array = res
	mat.m = a.m
	mat.n = b.n
	return mat
}

// LPN <a,b>
func (mat *Matrix) Vectproduct(r, s *Bitsvect) *Matrix {
	m := r.n
	res := make([]*Bitsvect, m)
	mat = NewMatrix(res)
	return mat
}

// Print a bits matrix
func (mat *Matrix) Print(a *Matrix) string {
	n := a.n // columns
	m := a.m // rows
	buf := make([][]byte, m)
	for i := 0; i < n; i++ {
		b := a.array[i]
		m1 := len(b.array)
		// first m1-1 elements
		for j := 0; j < m1-1; j++ {
			c := b.array[j]
			for k := 0; k < bitsPerUnit; k++ {
				d := uint(1 << k)
				s := c & d
				if s == d {
					buf[j*bitsPerUnit+k] = append(buf[j*bitsPerUnit+k], '1')
				} else {
					buf[j*bitsPerUnit+k] = append(buf[j*bitsPerUnit+k], '0')
				}
			}
		}
		// the last element
		l := m1 - 1
		m2 := m - l*bitsPerUnit
		c := b.array[l]
		for k := 0; k < m2; k++ {
			s := c & 1
			if s == 1 {
				buf[l*bitsPerUnit+k] = append(buf[l*bitsPerUnit+k], '1')
			} else {
				buf[l*bitsPerUnit+k] = append(buf[l*bitsPerUnit+k], '0')
			}
			c >>= 1
		}
	}
	var res string
	for i := 0; i < len(buf); i++ {
		res += string(buf[i]) + "\n"
	}
	return res
}

// mat=(a1,a2,\cdots,an)
func NewMatrix(a []*Bitsvect) *Matrix {
	mat := new(Matrix)
	mat.n = len(a)
	mat.array = a
	mat.m = a[0].n
	return mat
}
