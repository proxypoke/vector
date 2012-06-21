package vector

import (
	"math"
)

// A vector over float64 values
type Vector struct {
	dims []float64 // the elements of the vector
	ndim uint      // the dimension of the vector
}

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
