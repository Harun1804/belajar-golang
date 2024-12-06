package geometry

import "math"

type Lingkaran struct {
	Diameter float64
}

func (l Lingkaran) JariJari() float64 {
	return l.Diameter / 2
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * math.Pow(l.JariJari(), 2)
}

func (l Lingkaran) Keliling() float64 {
	return math.Pi * l.Diameter
}