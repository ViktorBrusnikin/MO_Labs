package dichotomy

import (
	"testgo/result"
	"testgo/vectornd"
)

func Dichotomy(
	result *result.Result,
	function func(vectornd.VectorND) float64,
	lhs vectornd.VectorND,
	rhs vectornd.VectorND,
	eps float64,
	maxIterations int,
) *result.Result {

	result.Clear()

	result.Dim = lhs.Dim()

	dir := rhs.SubVec(lhs).DivConst(rhs.SubVec(lhs).Length()).MulConst(eps * 0.1)

	i := 0
	for ; rhs.SubVec(lhs).Length() > 2*eps && i < maxIterations; i++ {
		x_c := lhs.AddVec(rhs).MulConst(0.5)
		x_l := x_c.SubVec(dir)
		x_r := x_c.AddVec(dir)

		if function(x_l) > function(x_r) {
			lhs = x_l
		} else {
			rhs = x_r
		}
	}

	result.Result = lhs.AddVec(rhs).MulConst(0.5)
	result.Iterations = i
	result.NumberMeasurements = i * 2
	result.Accuracy = rhs.SubVec(lhs).Length() * 0.5

	return result
}
