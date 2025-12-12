package vector3d

import "math"

type Vector3D struct {
	X float64
	Y float64
	Z float64
}

func NewVector3D(x, y, z float64) Vector3D {
	return Vector3D{
		X: x,
		Y: y,
		Z: z,
	}
}

func NewZeroVector3D() Vector3D {
	return NewVector3D(0.0, 0.0, 0.0)
}

func (v Vector3D) AddVec(other Vector3D) Vector3D {
	return Vector3D{
		v.X + other.X,
		v.Y + other.Y,
		v.Z + other.Z,
	}
}

func (v Vector3D) AddConst(C float64) Vector3D {
	return Vector3D{
		v.X + C,
		v.Y + C,
		v.Z + C,
	}
}

func (v Vector3D) SubVec(other Vector3D) Vector3D {
	return Vector3D{
		v.X - other.X,
		v.Y - other.Y,
		v.Z - other.Z,
	}
}

func (v Vector3D) SubConst(C float64) Vector3D {
	return Vector3D{
		v.X - C,
		v.Y - C,
		v.Z - C,
	}
}

func (v Vector3D) MulConst(C float64) Vector3D {
	return Vector3D{
		v.X * C,
		v.Y * C,
		v.Z * C,
	}
}

func (v Vector3D) DivConst(C float64) Vector3D {
	return Vector3D{
		v.X / C,
		v.Y / C,
		v.Z / C,
	}
}

func (v Vector3D) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector3D) DotProduct(other Vector3D) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}
