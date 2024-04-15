package mechanics

import (
	"math"
)

type Coorder interface {
	GetX() int32
	GetY() int32
}

func Distance(a Coorder, b Coorder) float64 {
	return math.Sqrt(math.Pow(float64(b.GetX())-float64(a.GetX()), 2) +
		math.Pow(float64(b.GetY())-float64(a.GetY()), 2))
}

func DistanceFromCenter(a Coorder) float64 {
	return math.Sqrt(math.Pow(-float64(a.GetX()), 2) +
		math.Pow(-float64(a.GetY()), 2))
}
