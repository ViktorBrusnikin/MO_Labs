package goldenratio

import (
	"testgo/result"
	"testgo/vectornd"
)

const PSI = 0.61803398874989484820
const PHI = 1.61803398874989484820

func GoldenRatio(
	result *result.Result,
	function func(vectornd.VectorND) float64,
	lhs vectornd.VectorND,
	rhs vectornd.VectorND,
	eps float64,
	maxIterations int,
) *result.Result {

	result.Clear()

	result.Dim = lhs.Dim()

	x_r := lhs.AddVec((rhs.SubVec(lhs)).MulConst(PSI))
	x_l := rhs.SubVec((rhs.SubVec(lhs)).MulConst(PSI))

	f_l := function(x_l)
	f_r := function(x_r)

	i := 0
	for ; rhs.SubVec(lhs).Length() > 2*eps && i < maxIterations; i++ {
		if f_l > f_r {
			lhs = x_l
			x_l = x_r
			f_l = f_r
			x_r = lhs.AddVec((rhs.SubVec(lhs)).MulConst(PSI))
			f_r = function(x_r)
		} else {
			rhs = x_r
			x_r = x_l
			f_r = f_l
			x_l = rhs.SubVec((rhs.SubVec(lhs)).MulConst(PSI))
			f_l = function(x_l)
		}
	}

	result.Iterations = i
	result.NumberMeasurements = i + 2
	result.Result = rhs.AddVec(lhs).MulConst(0.5)
	result.Accuracy = rhs.SubVec(lhs).Length() * 0.5

	return result
}
