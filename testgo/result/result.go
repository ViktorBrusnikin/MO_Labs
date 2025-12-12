package result

import (
	"fmt"
	"testgo/vectornd"
)

type Result struct {
	Result             vectornd.VectorND
	NumberMeasurements int
	Iterations         int
	Accuracy           float64
	Dim                int
}

func NewResult() *Result {
	return &Result{}
}

func (r *Result) Clear() {
	r.Result = vectornd.NewZeroVectorND(r.Dim)
	r.NumberMeasurements = 0
	r.Iterations = 0
	r.Accuracy = 0
	r.Dim = 0
}

func (res *Result) String() string {
	return fmt.Sprintf(
		"\tРезультат: %f\n\tКоличество измерений: %d\n\tТочность: %e\n\tКоличество итераций: %d\n",
		res.Result, res.NumberMeasurements, res.Accuracy, res.Iterations,
	)
}
