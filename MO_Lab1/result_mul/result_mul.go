package result_mul

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type ResultMul struct {
	Result            mat.VecDense
	CountMeasurements int
	Accuracy          float64
	Iterations        int
}

func New(result *mat.VecDense, countMeasurements int, accuracy float64, iterations int) *ResultMul {
	return &ResultMul{
		Result:            *result, // копируем значение из указателя
		CountMeasurements: countMeasurements,
		Accuracy:          accuracy,
		Iterations:        iterations,
	}
}

func (res *ResultMul) Clear() {
	res.Result = *mat.NewVecDense(3, []float64{0.0, 0.0, 0.0})
	res.CountMeasurements = 0
	res.Accuracy = 0.0
	res.Iterations = 0
}

func (res *ResultMul) String() string {
	return fmt.Sprintf(
		"\tРезультат: %f\n\tКоличество измерений: %d\n\tТочность: %e\n\tКоличество итераций: %d\n",
		res.Result, res.CountMeasurements, res.Accuracy, res.Iterations,
	)
}
