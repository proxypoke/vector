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
		err = IndexError(n)
		return
	}
	val = v.dims[n]
	return
}

// Set the value of the nth element in the Vector.
func (v Vector) Set(n uint, x float64) (err error) {
	if n >= v.Dim() {
		err = IndexError(n)
		return
	}
	v.dims[n] = x
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
func (v Vector) Add(other Vector) (Vector, error) {
	err := checkDims(v, other)
	if err == nil {
		for i := range v.dims {
			v.dims[i] += other.dims[i]
		}
	}
	return v, err
}

// Substract another Vector, in-place.
func (v Vector) Substract(other Vector) (Vector, error) {
	err := checkDims(v, other)
	if err == nil {
		for i := range v.dims {
			v.dims[i] -= other.dims[i]
		}
	}
	return v, err
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

// Cross product with another vector, in-place.
// Returns error when ndim of either vector != 3.
func (v Vector) CrossProduct(other Vector) (Vector, error) {
	if v.ndim != 3 || other.ndim != 3 {
		err := CrossError{v.ndim, other.ndim}
		return New(0), err
	}
	x := v.dims[1]*other.dims[2] - v.dims[2]*other.dims[1]
	y := v.dims[2]*other.dims[0] - v.dims[0]*other.dims[2]
	z := v.dims[0]*other.dims[1] - v.dims[1]*other.dims[0]
	v.dims[0] = x
	v.dims[1] = y
	v.dims[2] = z
	return v, nil
}

// ============================== [ Functions ] ===============================

// Check if two vectors have the same dimension.
func checkDims(a, b Vector) (err error) {
	if a.ndim != b.ndim {
		err = DimError{a.ndim, b.ndim}
	}
	return
}

// Add two Vectors, returning a new Vector.
func Add(a, b Vector) (Vector, error) {
	return a.Copy().Add(b)
}

// Substract two Vectors, returning new Vector.
func Substract(a, b Vector) (Vector, error) {
	return a.Copy().Substract(b)
}

// Scalar multiplication of a Vector, returning a new Vector.
func Scale(v Vector, x float64) Vector {
	return v.Copy().Scale(x)
}

// Normalize a vector, returning a new Vector.
func Normalize(v Vector) Vector {
	return v.Copy().Normalize()
}

// Dot-product of two Vectors.
func DotProduct(a, b Vector) (dot float64, err error) {
	for i := range a.dims {
		dot += a.dims[i] * b.dims[i]
	}
	return
}

// Angle (theta) between two vectors.
func Angle(a, b Vector) (theta float64, err error) {
	err = checkDims(a, b)
	if err == nil {
		norm_a := Normalize(a)
		norm_b := Normalize(b)
		dot, _ := DotProduct(norm_a, norm_b)
		theta = math.Acos(dot)
	}
	return
}

// Cross product of two vectors.
// Returns error when ndim of either vector != 3.
func CrossProduct(a, b Vector) (Vector, error) {
	return a.Copy().CrossProduct(b)
}
