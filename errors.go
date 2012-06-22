// Copywrong 2012 slowpoke <proxypoke at lavabit dot com>
// Repository: https://github.com/proxypoke/vector.go
//
// This program is free software under the terms of the
// Do What The Fuck You Want To Public License v2,
// which can be found in a file called COPYING included
// with this program or at http://sam.zoy.org/wtfpl/COPYING

package vector

import (
	"strconv"
)

// For missmatched dimensions.
type DimError struct {
	Dim_a, Dim_b uint
}

type (
	IndexError uint     // For out-of-range indexes.
	CrossError DimError // For cross products where either ndim != 3.

)

func (e DimError) Error() string {
	return "Missmatched dimensions: " +
		strconv.Itoa(int(e.Dim_b)) +
		" != " +
		strconv.Itoa(int(e.Dim_b))
}

func (e IndexError) Error() string {
	return "Index out of range: " + strconv.Itoa(int(e))
}

func (e CrossError) Error() string {
	return "Invalid dimensions: " +
		strconv.Itoa(int(e.Dim_a)) +
		", " +
		strconv.Itoa(int(e.Dim_b)) +
		" (must be 3)"
}
