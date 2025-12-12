package main

import (
	"MO_LAB1/multi_dim_methods/dichotomy_mul"
	"MO_LAB1/one_dim_methods/dichotomy"
	"MO_LAB1/one_dim_methods/fibonacci"
	"MO_LAB1/one_dim_methods/goldenratio"
	"MO_LAB1/result"
	"MO_LAB1/result_mul"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

const PSI = 0.61803398874989484820

func main() {

	lhs := 0.0
	rhs := 5.0
	eps := 1e-6
	maxIteration := 51

	test_func_one_dim := func(x float64) float64 {
		return (x - 2) * (x - 5)
	}

	test_func_mul_dim := func(x mat.VecDense) float64 {
		// Создаём векторы нужной длины
		first_part := mat.NewVecDense(x.Len(), nil)
		second_part := mat.NewVecDense(x.Len(), nil)

		// Заполняем их
		for i := 0; i < x.Len(); i++ {
			first_part.SetVec(i, x.AtVec(i)-2.0)
			second_part.SetVec(i, x.AtVec(i)-5.0)
		}

		// Скалярное произведение
		return mat.Dot(first_part, second_part)

	}

	res := result.New(0.0, 0, 0.0, 0)

	fmt.Println("~~~~~~~~~~~~~Одномерные методы~~~~~~~~~~~~~")
	dichotomy.Dichotomy(res, test_func_one_dim, lhs, rhs, eps, maxIteration)
	fmt.Println("-------------------------------------------")
	fmt.Println("Результат Дихотомии:\n", res)
	fmt.Println("-------------------------------------------")

	goldenratio.GoldenRatio(res, test_func_one_dim, lhs, rhs, eps*PSI, maxIteration)
	fmt.Println("-------------------------------------------")
	fmt.Println("Результат Золотого сечения:\n", res)
	fmt.Println("-------------------------------------------")

	fibonacci.Fibonacci(res, test_func_one_dim, lhs, rhs, eps)
	fmt.Println("-------------------------------------------")
	fmt.Println("Результат Фибоначчи:\n", res)
	fmt.Println("-------------------------------------------")

	resMul := result_mul.New(mat.NewVecDense(3, nil), 0, 0.0, 0)
	lhs_mul := mat.NewVecDense(3, []float64{0.0, 0.0, 0.0})
	rhs_mul := mat.NewVecDense(3, []float64{5.0, 5.0, 5.0})

	fmt.Println("~~~~~~~~~~~~~Многомерные методы~~~~~~~~~~~~~")
	dichotomy_mul.DichotomyMul(resMul, test_func_mul_dim, *lhs_mul, *rhs_mul, eps, maxIteration)
	fmt.Println("-------------------------------------------")
	fmt.Println("Результат Дихотомии:\n", res)
	fmt.Println("-------------------------------------------")
}
