package vector

import "math"

type Vec []float64

func (a Vec) Plus(b Vec) (c Vec) {
	for k, it := range a {
		c = append(c, it+b[k])
	}
	return
}

func (a Vec) Minus(b Vec) (c Vec) {
	for k, it := range a {
		c = append(c, it-b[k])
	}
	return
}

func (a Vec) ScalarMultiply(n float64) (c Vec) {
	for _, it := range a {
		c = append(c, it*n)
	}
	return
}

func (a Vec) Magnitude() float64 {
	sumOfSquare := 0.0
	for _, it := range a {
		sumOfSquare += it * it
	}
	sqrt := math.Sqrt(sumOfSquare)
	if sqrt < 0 {
		sqrt *= -1
	}
	return sqrt
}

func (a Vec) Normalize() (c Vec) {
	scalar := 1 / a.Magnitude()
	return a.ScalarMultiply(scalar)
}
