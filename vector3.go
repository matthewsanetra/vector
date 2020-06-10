package vector

import (
	"fmt"
	"math"
)

type Vector3f struct {
	X, Y, Z float64
}

func (v *Vector3f) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector3f) String() string {
	return fmt.Sprintf("<%.6f, %.6f, %.6f>|%.6f|", v.X, v.Y, v.Z, v.Mag())
}

func Add3f(v1, v2, out *Vector3f) {
	out.X = v1.X + v2.X
	out.Y = v1.Y + v2.Y
	out.Z = v1.Z + v2.Z
}

func Sub3f(v1, v2, out *Vector3f) {
	out.X = v1.X - v2.X
	out.Y = v1.Y - v2.Y
	out.Z = v1.Z - v2.Z
}

func Mul3f(v *Vector3f, k float64, out *Vector3f) {
	out.X = v.X * k
	out.Y = v.Y * k
	out.Z = v.Z * k
}

func Div3f(v *Vector3f, k float64, out *Vector3f) {
	out.X = v.X / k
	out.Y = v.Y / k
	out.Z = v.Z / k
}
