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

func (a Vec) DotProduct(b Vec) (sum float64) {
	for k, it := range a {
		sum += it * b[k]
	}
	return
}

func (a Vec) AngleRad(b Vec) float64 {
	return math.Acos(a.DotProduct(b) / (a.Magnitude() * b.Magnitude()))
}

func (a Vec) AngleDegrees(b Vec) float64 {
	return a.AngleRad(b) * 180 / math.Pi
}

func (a Vec) IsParalel(b Vec) bool  {
	aMagnitude := a.Magnitude()
	bMagnitude := b.Magnitude()
	if bMagnitude > aMagnitude {
		bMagnitude, aMagnitude = aMagnitude, bMagnitude
	}
	scalar := aMagnitude / bMagnitude
	return scalar == float64(int64(scalar))
}

const (
	tolerance = 0.00000000001
)

func (a Vec) IsOrthogonal(b Vec) bool {
	return math.Abs(a.DotProduct(b)) < tolerance
}
