package dichotomy_mul

import (
	"MO_LAB1/result_mul"

	"gonum.org/v1/gonum/mat"
)

func DichotomyMul(result_func *result_mul.ResultMul, function func(mat.VecDense) float64, lhs mat.VecDense, rhs mat.VecDense, eps float64, maxIteration int) *result_mul.ResultMul {
	var x_c_temp mat.VecDense
	var x_c mat.VecDense
	var x_l mat.VecDense
	var x_r mat.VecDense
	var f_l float64
	var f_r float64

	result_func.Clear()

	// diff = rhs - lhs
	var diff mat.VecDense
	diff.SubVec(&rhs, &lhs)

	// norm = ||diff||_2
	norm := mat.Norm(&diff, 2)

	// temp_dir = diff / norm
	var temp_dir mat.VecDense
	temp_dir.ScaleVec(1/norm, &diff)

	// dir = eps * temp_dir
	var dir mat.VecDense
	dir.ScaleVec(eps, &temp_dir)

	i := 0
	for ; i < maxIteration && mat.Norm(&diff, 2) > 2*eps; i++ {
		diff.SubVec(&rhs, &lhs)
		x_c_temp.AddVec(&rhs, &lhs)
		x_c.ScaleVec(0.5, &x_c_temp)

		x_r.AddVec(&x_r, &x_c)
		x_l.SubVec(&x_l, &x_c)
		f_l = function(x_l)
		f_r = function(x_r)

		if f_l > f_r {
			lhs = x_l
		} else {
			rhs = x_r
		}
	}
	result := mat.NewVecDense(3, nil)
	result.AddVec(&rhs, &lhs)
	result.ScaleVec(0.5, result)

	accuracy_temp := mat.NewVecDense(3, nil)
	accuracy_temp.SubVec(&rhs, &lhs)
	accuracy := mat.Norm(accuracy_temp, 2)

	result_func.Iterations = i
	result_func.CountMeasurements = i * 2
	result_func.Result = *result
	result_func.Accuracy = accuracy * 0.5

	return result_func
}
