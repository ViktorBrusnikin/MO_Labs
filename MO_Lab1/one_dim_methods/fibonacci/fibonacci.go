package fibonacci

import (
	"MO_LAB1/result"
	"math"
)

func Fibonacci(result_func *result.Result, function func(float64) float64, lhs float64, rhs float64, eps float64) *result.Result {
	result_func.Clear()
	condition := (rhs - lhs) / eps

	L_t := 0.0
	L_1 := 1.0
	L_2 := 1.0

	n := 0
	for ; L_2 < condition; n++ {
		L_t = L_1
		L_1 = L_2
		L_2 += L_t
	}

	x_l := lhs + ((L_2-L_1)/L_2)*(rhs-lhs)
	x_r := lhs + (L_1/L_2)*(rhs-lhs)
	f_l := function(x_l)
	f_r := function(x_r)

	result_func.CountMeasurements = n + 2
	result_func.Iterations = n
	for ; n > 0; n-- {
		L_t = L_2 - L_1
		L_2 = L_1
		L_1 = L_t

		if f_l > f_r {
			lhs = x_l
			x_l = x_r
			f_l = f_r
			x_r = lhs + (L_1/L_2)*(rhs-lhs)
			delta := (rhs - lhs) * 0.01
			if math.Abs(x_r-x_l) < delta {
				x_r += delta
			}
			f_r = function(x_r)
		} else {
			rhs = x_r
			x_r = x_l
			f_r = f_l
			x_l = lhs + ((L_2-L_1)/L_2)*(rhs-lhs)
			delta := (rhs - lhs) * 0.01
			if math.Abs(x_r-x_l) < delta {
				x_l -= delta
			}
			f_l = function(x_l)
		}
	}

	result_func.Result = (rhs + lhs) * 0.5
	result_func.Accuracy = math.Abs(rhs-lhs) * 0.5

	return result_func
}
