package geometry

import "math"

type Persegi struct {
	Sisi float64
}

func (p Persegi) Luas() float64 {
	return math.Pow(p.Sisi, 2)
}

func (p Persegi) Keliling() float64  {
	return p.Sisi * 4
}
