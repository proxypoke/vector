// Copywrong 2012 slowpoke <proxypoke at lavabit dot com>
// Repository: https://github.com/proxypoke/vector.go
//
// This program is free software under the terms of the
// Do What The Fuck You Want To Public License v2,
// which can be found in a file called COPYING included
// with this program or at http://sam.zoy.org/wtfpl/COPYING

package vector

import (
	"math/rand"
	"testing"
)

// Creates vectors with dimension from 0 to 99, checks if they actually have
// that dimension, then checks if the values are correctly initialized to 0.
func TestNew(t *testing.T) {
	var i, j uint
	for i = 0; i < 100; i++ {
		v := New(i)
		if v.Dim() != i {
			t.Errorf("Wrong dimension. Got %d, expected %d.", v.Dim(), i)
		}
		for j = 0; j < i; j++ {
			// XXX: If the Get method errors, this test will still pass. This
			// is because Get() would then return an uninitialized float64 for
			// val, which is 0 and therefore what the test expects.
			val, _ := v.Get(j)
			if val != 0 {
				t.Error("Newly initialized vector has a value != 0.")
			}
		}
	}
}

// Creates vectors with randomized slices, then checks whether they have the
// correct dimension (len(slice)) and whether they have been correctly
// initialized.
func TestNewFrom(t *testing.T) {
	var i, j uint
	for i = 0; i < 100; i++ {
		randslice := makeRandSlice(i)
		v := NewFrom(randslice)
		if v.Dim() != i {
			t.Errorf("Wrong dimension. Got %d, expected %d.", v.Dim(), i)
		}
		for j = 0; j < i; j++ {
			val, _ := v.Get(j)
			if val != randslice[j] {
				t.Error(
					"Wrong values in vector initialized from a random slice.")
			}
		}
	}
}

// Creates pseudo-random vectors with various dimensions, copies them and
// verifies that the new vector is equal.
func TestCopy(t *testing.T) {
	var i uint
	for i = 0; i < 100; i++ {
		v := makeRandomVector(i)
		w := v.Copy()
		for j := range v.dims {
			if v.dims[j] != w.dims[j] {
				t.Error("Copied vector has differnt value.")
			}
		}
	}
}

// Creates pseudo-random vectors with various dimensions, then check if Get()
// returns the correct values and errors on out-of-range indexes.
func TestGet(t *testing.T) {
	var i uint
	for i = 0; i < 100; i++ {
		v := makeRandomVector(i)
		for j, val := range v.dims {
			getval, err := v.Get(uint(j))
			if err != nil {
				t.Error("Get() errored on a correct index.")
			}
			if val != getval {
				t.Error("Get() returned a wrong value.")
			}
		}
		_, err := v.Get(v.Dim())
		if err == nil {
			t.Error("Get didn't error on an out-of-range index.")
		}
	}
}

// Creates uninitialized vectors of various dimensions, then sets their values
// to pseudo-random values. It then compares those values to check if they
// were set correctly. Also verifies is Set() correctly errors on out-of-range
// indexes.
func TestSet(t *testing.T) {
	var i, j uint
	for i = 0; i < 100; i++ {
		v := New(i)
		for j = 0; j < i; j++ {
			val := rand.ExpFloat64()
			err := v.Set(j, val)
			if err != nil {
				t.Error("Set() errored on a correct index.")
			}
			if v.dims[j] != val {
				t.Error("Set didn't correctly set a value.")
			}
		}
		err := v.Set(v.Dim(), 0)
		if err == nil {
			t.Error("Set didn't error on an out-of-range index.")
		}
	}
}

// Creates a vector with known length, then compares the expected value with
// what Len() returns.
func TestLen(t *testing.T) {
	v := New(1)
	v.Set(0, 2)	// has length 2
	if v.Len() != 2 {
		t.Error("Len returned a wrong length")
	}
}

// Helper function, makes pseudo-random slices.
func makeRandSlice(length uint) (randslice []float64) {
	randslice = make([]float64, length)
	for i := range randslice {
		randslice[i] = rand.ExpFloat64()
	}
	return
}

// Helper function, make a pseudo-random Vector with dimension dim.
func makeRandomVector(dim uint) Vector {
	return NewFrom(makeRandSlice(dim))
}
