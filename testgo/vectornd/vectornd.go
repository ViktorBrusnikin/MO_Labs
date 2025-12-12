package vectornd

import (
	"fmt"
	"math"
)

type VectorND struct {
	coords []float64
}

func NewVectorND(coords []float64) VectorND {
	c := make([]float64, len(coords))
	copy(c, coords)
	return VectorND{coords: c}
}

func NewZeroVectorND(dim int) VectorND {
	return VectorND{coords: make([]float64, dim)}
}

func (v VectorND) Dim() int {
	return len(v.coords)
}

func (v VectorND) Coords() []float64 {
	c := make([]float64, len(v.coords))
	copy(c, v.coords)
	return c
}

func (v VectorND) String() string {
	return fmt.Sprintf("VectorND%v", v.coords)
}

func (v VectorND) AddVec(other VectorND) VectorND {
	if v.Dim() != other.Dim() {
		panic("dimension mismatch in AddVec")
	}
	result := make([]float64, v.Dim())
	for i := range v.coords {
		result[i] = v.coords[i] + other.coords[i]
	}
	return VectorND{coords: result}
}

func (v VectorND) AddConst(c float64) VectorND {
	result := make([]float64, v.Dim())
	for i := range v.coords {
		result[i] = v.coords[i] + c
	}
	return VectorND{coords: result}
}

func (v VectorND) SubVec(other VectorND) VectorND {
	if v.Dim() != other.Dim() {
		panic("dimension mismatch in SubVec")
	}
	result := make([]float64, v.Dim())
	for i := range v.coords {
		result[i] = v.coords[i] - other.coords[i]
	}
	return VectorND{coords: result}
}

func (v VectorND) SubConst(c float64) VectorND {
	result := make([]float64, v.Dim())
	for i := range v.coords {
		result[i] = v.coords[i] - c
	}
	return VectorND{coords: result}
}

func (v VectorND) MulConst(c float64) VectorND {
	result := make([]float64, v.Dim())
	for i := range v.coords {
		result[i] = v.coords[i] * c
	}
	return VectorND{coords: result}
}

func (v VectorND) DivConst(c float64) VectorND {
	if c == 0 {
		panic("division by zero in DivConst")
	}
	result := make([]float64, v.Dim())
	for i := range v.coords {
		result[i] = v.coords[i] / c
	}
	return VectorND{coords: result}
}

func (v VectorND) Length() float64 {
	sum := 0.0
	for i := range v.coords {
		sum += v.coords[i] * v.coords[i]
	}
	return math.Sqrt(sum)
}

func (v VectorND) DotProduct(other VectorND) float64 {
	if v.Dim() != other.Dim() {
		panic("dimension mismatch in DotProduct")
	}
	sum := 0.0
	for i := range v.coords {
		sum += v.coords[i] * other.coords[i]
	}
	return sum
}
