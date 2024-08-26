package points

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y)
}

func Abs(p Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// or we can create a method(an instance of the struct). This is also a function
func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// method with argument
func (p Point) CompateTo(other Point) Point {
	if p.Abs() > other.Abs() {
		return p
	}
	return other
}
