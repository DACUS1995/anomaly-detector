package detector

import (
	"fmt"
	"log"
	"testing"

	"gonum.org/v1/gonum/mat"
)

//http://gregorygundersen.com/blog/2019/10/30/scipy-multivariate/
//https://godoc.org/gonum.org/v1/gonum
func TestCovarianceInverse(t *testing.T) {
	a := mat.NewDense(2, 2, []float64{
		4, 0,
		0, 4,
	})

	b := mat.NewVecDense(2, []float64{
		2, 2,
	})
	var x mat.VecDense
	err := x.SolveVec(a, b)
	if err != nil {
		log.Fatalf("no solution: %v", err)
	}

	c := mat.NewVecDense(2, []float64{
		2, 2,
	})
	var y mat.VecDense
	y.MulElemVec(&x, c)

	// Print the result using the formatter.
	fx := mat.Formatted(&y, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("x = %v", fx)
}
