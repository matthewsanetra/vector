package vector

import (
	"fmt"
	"math"
)

type Vector2f struct {
	X, Y float64
}

func (v *Vector2f) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector2f) String() string {
	return fmt.Sprintf("<%.6f, %.6f>|%.6f|", v.X, v.Y, v.Mag())
}

func Add2f(v1, v2, out *Vector2f) {
	out.X = v1.X + v2.X
	out.Y = v1.Y + v2.Y
}

func Sub2f(v1, v2, out *Vector2f) {
	out.X = v1.X - v2.X
	out.Y = v1.Y - v2.Y
}

func Mul2f(v *Vector2f, k float64, out *Vector2f) {
	out.X = v.X * k
	out.Y = v.Y * k
}

func Div2f(v *Vector2f, k float64, out *Vector2f) {
	out.X = v.X / k
	out.Y = v.Y / k
}
