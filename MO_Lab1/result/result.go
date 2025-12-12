package result

import "fmt"

type Result struct {
	Result            float64
	CountMeasurements int
	Accuracy          float64
	Iterations        int
}

func New(result float64, countMeasurements int, accuracy float64, iterations int) *Result {
	return &Result{
		Result:            result,
		CountMeasurements: countMeasurements,
		Accuracy:          accuracy,
		Iterations:        iterations,
	}
}

func (res *Result) Clear() {
	res.Result = 0.0
	res.CountMeasurements = 0
	res.Accuracy = 0.0
	res.Iterations = 0
}

func (res *Result) String() string {
	return fmt.Sprintf(
		"\tРезультат: %f\n\tКоличество измерений: %d\n\tТочность: %e\n\tКоличество итераций: %d\n",
		res.Result, res.CountMeasurements, res.Accuracy, res.Iterations,
	)
}
