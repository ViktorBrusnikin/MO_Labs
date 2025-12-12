package main

import (
	"fmt"
	"testgo/multidimensional_methods/coordinatedescent"
	"testgo/multidimensional_methods/dichotomy"
	"testgo/multidimensional_methods/fibonacci"
	"testgo/multidimensional_methods/goldenratio"
	"testgo/result"
	"testgo/vectornd"
)

func call_methods(
	result *result.Result,
	test_func func(vectornd.VectorND) float64,
	lhs vectornd.VectorND,
	rhs vectornd.VectorND,
	eps float64,
	maxIterations int,
	lambda float64,
) {
	fmt.Println("Многомерные методы")

	fmt.Println("Результат работы Дихотомии")
	dichotomy.Dichotomy(result, test_func, lhs, rhs, eps, maxIterations)
	fmt.Println(result)

	fmt.Println("Результат работы Золотого сечения")
	goldenratio.GoldenRatio(result, test_func, lhs, rhs, eps*goldenratio.PSI, maxIterations)
	fmt.Println(result)

	fmt.Println("Результат работы Фибоначчи")
	fibonacci.Fibonacci(result, test_func, lhs, rhs, eps)
	fmt.Println(result)

	fmt.Println("Результат работы Покоординатного спуска")
	coordinatedescent.CoorDinatedescent3D(result, test_func, lhs, eps, lambda, maxIterations)
	fmt.Println(result)
}

func main() {
	test_func := func(vec vectornd.VectorND) float64 {
		return vec.DotProduct(vec.SubConst(1))
	}
	lhs := vectornd.NewVectorND([]float64{0.0, 0.0})
	rhs := vectornd.NewVectorND([]float64{1.0, 1.0})
	result := result.NewResult()
	eps := 1e-7
	maxIterations := 1000
	lambda := 0.1

	call_methods(result, test_func, lhs, rhs, eps, maxIterations, lambda)
}
