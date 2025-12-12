package coordinatedescent

import (
	"testgo/multidimensional_methods/fibonacci"
	"testgo/result"
	"testgo/vectornd"
)

func CoorDinatedescent3D(
	result *result.Result,
	function func(vectornd.VectorND) float64,
	x_0 vectornd.VectorND,
	eps float64,
	lambda float64,
	maxIteration int,
) *result.Result {

	result.Clear()

	n := x_0.Dim()

	result.Dim = n

	e_ort := make([]vectornd.VectorND, n)

	for i := 0; i < n; i++ {
		coords := make([]float64, n)
		coords[i] = 1.0
		e_ort[i] = vectornd.NewVectorND(coords)
	}

	var number_measurements int

	x_cur := x_0
	var x_next vectornd.VectorND
	break_criteria := 0
	i := 0
	for ; i < maxIteration; i++ {

		index := i % n

		x_l := x_cur.SubVec(e_ort[index].MulConst(eps))
		x_r := x_cur.AddVec(e_ort[index].MulConst(eps))

		f_l := function(x_l)
		f_r := function(x_r)

		if f_l > f_r {
			res := fibonacci.Fibonacci(result, function, x_cur, x_cur.AddVec(e_ort[index].MulConst(lambda)), eps)
			x_next = res.Result
			number_measurements += res.NumberMeasurements
		} else {
			res := fibonacci.Fibonacci(result, function, x_cur, x_cur.SubVec(e_ort[index].MulConst(lambda)), eps)
			x_next = res.Result
			number_measurements += res.NumberMeasurements
		}

		result.Clear()

		if x_next.SubVec(x_cur).Length() < 2*eps {
			break_criteria += 1
			if break_criteria == n {
				break
			}
		} else {
			break_criteria = 0
		}

		x_cur = x_next

	}

	result.Iterations = i
	result.NumberMeasurements = number_measurements
	result.Result = x_next
	result.Accuracy = x_next.SubVec(x_cur).Length()

	return result
}
