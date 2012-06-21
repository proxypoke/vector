// Copywrong 2012 slowpoke <proxypoke at lavabit dot com>
// Repository: https://github.com/proxypoke/vector.go
//
// This program is free software under the terms of the
// Do What The Fuck You Want To Public License v2,
// which can be found in a file called COPYING included
// with this program or at http://sam.zoy.org/wtfpl/COPYING

package vector

func Map(f func(float64) float64, seq []float64) (newseq []float64) {
	for _, val := range seq {
		newseq = append(newseq, f(val))
	}
	return
}

func Reduce(f func(float64, float64) float64, seq []float64) (result float64) {
	result = f(seq[0], seq[1])
	for n := 2; n < len(seq); n++ {
		result = f(result, seq[n])
	}
	return
}
