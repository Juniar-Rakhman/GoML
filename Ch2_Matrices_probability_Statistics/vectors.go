package main

import (
	"fmt"

	"github.com/gonum/blas/blas64"
	"github.com/gonum/floats"
	"github.com/gonum/matrix/mat64"
)

func main() {
	vectorA := []float64{11.0, 5.2, -1.3}
	vectorB := []float64{-7.2, 4.2, 5.1}

	dotProduct := floats.Dot(vectorA, vectorB)
	fmt.Printf("The dot product of A and B is: %0.2f\n", dotProduct)

	floats.Scale(1.5, vectorA)
	fmt.Printf("Scaling A by 1.5 gives: %v\n", vectorA)

	// Compute the norm/length of B.
	normB := floats.Norm(vectorB, 2)
	fmt.Printf("The norm/length of B is: %0.2f\n", normB)

	vectorC := mat64.NewVector(3, []float64{11.0, 5.2, -1.3})
	vectorD := mat64.NewVector(3, []float64{-7.2, 4.2, 5.1})

	dotProduct2 := mat64.Dot(vectorC, vectorD)

	fmt.Printf("The dot product of C and D is: %0.2f\n", dotProduct2)

	// Scale each element of A by 1.5.
	vectorC.ScaleVec(1.5, vectorC)
	fmt.Printf("Scaling C by 1.5 gives: %v\n", vectorC)

	// Compute the norm/length of B.
	normD := blas64.Nrm2(3, vectorD.RawVector())
	fmt.Printf("The norm/length of B is: %0.2f\n", normD)
}
