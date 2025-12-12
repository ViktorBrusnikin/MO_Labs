package fibonacci

import (
	"testgo/result"
	"testgo/vectornd"
)

func Fibonacci(
	result *result.Result,
	function func(vectornd.VectorND) float64,
	lhs vectornd.VectorND,
	rhs vectornd.VectorND,
	eps float64,
) *result.Result {

	result.Clear()

	result.Dim = lhs.Dim()

	L_1, L_2 := 1.0, 1.0

	n := 0
	for ; L_1 < rhs.SubVec(lhs).Length()/eps; n++ {
		L_1, L_2 = L_1+L_2, L_1
	}

	x_r := lhs.AddVec(rhs.SubVec(lhs).MulConst(L_2 / L_1))
	x_l := lhs.AddVec(rhs.SubVec(lhs).MulConst((L_1 - L_2) / L_1))

	f_r := function(x_r)
	f_l := function(x_l)

	result.Iterations = n
	result.NumberMeasurements = n + 2

	for ; n > 0; n-- {
		L_1, L_2 = L_2, L_1-L_2

		if f_l > f_r {
			lhs = x_l
			x_l = x_r
			f_l = f_r
			x_r = lhs.AddVec(rhs.SubVec(lhs).MulConst(L_2 / L_1))
			delta := rhs.SubVec(lhs).Length() * 0.01
			if x_r.SubVec(x_l).Length() < delta {
				x_r = x_r.AddVec((rhs.SubVec(lhs).DivConst(rhs.SubVec(lhs).Length())).MulConst(delta))
			}
			f_r = function(x_r)
		} else {
			rhs = x_r
			x_r = x_l
			f_r = f_l
			x_l = lhs.AddVec(rhs.SubVec(lhs).MulConst((L_1 - L_2) / L_1))
			delta := rhs.SubVec(lhs).Length() * 0.01
			if x_r.SubVec(x_l).Length() < delta {
				x_l = x_l.AddVec((rhs.SubVec(lhs).DivConst(rhs.SubVec(lhs).Length())).MulConst(delta))
			}
			f_l = function(x_l)
		}
	}

	result.Result = rhs.AddVec(lhs).MulConst(0.5)
	result.Accuracy = rhs.SubVec(lhs).Length() * 0.5

	return result
}
