// Copywrong 2012 slowpoke <proxypoke at lavabit dot com>
// Repository: https://github.com/proxypoke/vector.go
//
// This program is free software under the terms of the
// Do What The Fuck You Want To Public License v2,
// which can be found in a file called COPYING included
// with this program or at http://sam.zoy.org/wtfpl/COPYING

package vector

import (
	"math"
)

// A vector over float64 values
type Vector struct {
	dims []float64 // the elements of the vector
	ndim uint      // the dimension of the vector
}

// ============================= [ Constructors ] =============================

// Create a Vector with dimension n, with all values initialized to 0.
func New(n uint) (v Vector) {
	v.dims = make([]float64, n)
	v.ndim = n
	return
}

// Create a Vector from a slice. Its dimension is equal to len(slice).
func NewFrom(dims []float64) (v Vector) {
	v.dims = dims
	v.ndim = uint(len(dims))
	return
}

// Make a deep copy of the Vector.
func (v Vector) Copy() Vector {
	return NewFrom(v.dims)
}

// =========================== [ General Methods ] ============================

// Get the dimension of the Vector.
func (v Vector) Dim() uint {
	return v.ndim
}

// Get the value of the nth element in the Vector.
func (v Vector) Get(n uint) (val float64, err error) {
	if n >= v.Dim() {
		return 0, IndexError(n)
	}
	val = v.dims[n]
	return
}

// Calculate the length of the Vector.
func (v Vector) Len() (result float64) {
	for _, val := range v.dims {
		result += math.Pow(val, 2)
	}
	return math.Sqrt(result)
}

// ========================= [ In-place operations ] ==========================

// Add another Vector, in-place.
func (v Vector) Add(other Vector) Vector {
	for i := range v.dims {
		v.dims[i] += other.dims[i]
	}
	return v
}

// Substract another Vector, in-place.
func (v Vector) Substract(other Vector) Vector {
	for i := range v.dims {
		v.dims[i] -= other.dims[i]
	}
	return v
}

// In-place scalar multiplication.
func (v Vector) Scale(x float64) Vector {
	for i := range v.dims {
		v.dims[i] *= x
	}
	return v
}

// Normalize the Vector (length == 1). In-place.
func (v Vector) Normalize() Vector {
	l := v.Len()
	for i := range v.dims {
		v.dims[i] /= l
	}
	return v
}

// ============================== [ Functions ] ===============================

// Dot-product of two Vectors.
func DotProduct(A, B Vector) (dot float64) {
	for i := range A.dims {
		dot += A.dims[i] * B.dims[i]
	}
	return
}
