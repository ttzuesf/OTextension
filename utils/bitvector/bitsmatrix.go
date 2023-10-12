package bitvector

type Matrix struct {
	array []Bitsvect
	m     int // rows
	n     int // columns
}

func (mat *Matrix) MatXor(a, b Matrix) *Matrix {
	if a.m != b.m && a.n != b.n {
		return nil
	}
	res := make([]Bitsvect, b.n)
	for i := 0; i < b.n; i++ {
		res[i] = Xor(a.array[i], b.array[i])
	}
	mat.array = res
	mat.m = a.m
	mat.n = b.n
	return mat
}

func (mat *Matrix) MatMul(r, s Bitsvect) *Matrix {
	m := len(r)
	r1 := make(Bitsvect, m)
	k := len(s)
	res := make([]Bitsvect, 0)
	for i := 0; i < k-1; i++ {
		a := s[i]
		for e := 0; e < bitsPerUnit; e++ {
			copy(r1, r)
			d := uint(1 << e)
			si := a & d //s[i]
			if si == d {
				for j := 0; j < m-1; j++ {
					for x := 0; x < bitsPerUnit; x++ {
						r1[j] ^= (1 << x)
					}
				}
				y := r1[m-1]
				x := 0
				for y != 0 {
					r1[m-1] ^= (1 << x)
					x++
					y >>= 1
				}
			}
			r2 := make(Bitsvect, m)
			copy(r2, r1)
			res = append(res, r2)
		}
	}
	a1 := s[k-1]
	for a1 != 0 {
		copy(r1, r)
		si := a1 & 1 //s[i]
		if si == 1 {
			for j := 0; j < m-1; j++ {
				for x := 0; x < bitsPerUnit; x++ {
					r1[j] ^= (1 << x)
				}
			}
			y := r1[m-1]
			e := 0
			for y != 0 {
				r1[m-1] ^= (1 << e)
				e++
				y >>= 1
			}
		}
		r2 := make(Bitsvect, m)
		copy(r2, r1)
		res = append(res, r2)
		a1 >>= 1
	}
	mat = NewMatrix(res)
	return mat
}

// Print a bits matrix
func (mat *Matrix) MatPrint(a *Matrix) string {
	n := a.n
	m := a.m
	buf := make([][]byte, m)
	for i := 0; i < n; i++ {
		b := a.array[i]
		m1 := len(a.array[i])
		// first m1-1 elements
		for j := 0; j < m1-1; j++ {
			c := b[j]
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
		c := b[l]
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

func NewMatrix(a []Bitsvect) *Matrix {
	mat := new(Matrix)
	mat.n = len(a)
	mat.array = make([]Bitsvect, mat.n)
	copy(mat.array, a)
	for i := 0; i < mat.n; i++ {
		t := Length(a[i])
		if mat.m < t {
			mat.m = t
		}
	}
	return mat
}
