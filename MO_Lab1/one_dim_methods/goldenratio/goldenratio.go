package goldenratio

import (
	"MO_LAB1/result"
	"math"
)

const PSI = 0.61803398874989484820

func GoldenRatio(result_func *result.Result, function func(float64) float64, lhs float64, rhs float64, eps float64, maxIteration int) *result.Result {

	result_func.Clear()
	x_l := rhs - PSI*(rhs-lhs)
	x_r := lhs + PSI*(rhs-lhs)
	f_l := function(x_l)
	f_r := function(x_r)

	i := 0
	for ; i < maxIteration && math.Abs(rhs-lhs) > 2*eps; i++ {
		if f_l > f_r {
			lhs = x_l
			x_l = x_r
			f_l = f_r
			x_r = lhs + PSI*(rhs-lhs)
			f_r = function(x_r)
		} else {
			rhs = x_r
			x_r = x_l
			f_r = f_l
			x_l = rhs - PSI*(rhs-lhs)
			f_l = function(x_l)
		}
	}

	result_func.Iterations = i
	result_func.Result = (rhs + lhs) * 0.5
	result_func.CountMeasurements = i + 2
	result_func.Accuracy = math.Abs(rhs-lhs) * 0.5

	return result_func
}
