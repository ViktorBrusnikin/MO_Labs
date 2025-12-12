package dichotomy

import (
	"MO_LAB1/result"
	"math"
)

func Dichotomy(result_func *result.Result, function func(float64) float64, lhs float64, rhs float64, eps float64, maxIteration int) *result.Result {

	result_func.Clear()
	i := 0
	for ; i < maxIteration && math.Abs(rhs-lhs) > 2*eps; i++ {
		x_c := (rhs + lhs) * 0.5
		x_r := x_c + eps*0.1
		x_l := x_c - eps*0.1
		f_l := function(x_l)
		f_r := function(x_r)

		if f_l > f_r {
			lhs = x_l
		} else {
			rhs = x_r
		}
	}

	result_func.Iterations = i
	result_func.CountMeasurements = i * 2
	result_func.Result = (rhs + lhs) * 0.5
	result_func.Accuracy = math.Abs(rhs-lhs) * 0.5

	return result_func
}
