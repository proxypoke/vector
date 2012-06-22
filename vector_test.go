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

// Helper function, makes pseudo-random slices.
func makeRandSlice(length uint) (randslice []float64) {
	randslice = make([]float64, length)
	for i := range randslice {
		randslice[i] = rand.ExpFloat64()
	}
	return
}
