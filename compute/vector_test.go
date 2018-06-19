package compute

import (
	"testing"
)

func TestMultiplyVector(t *testing.T) {
	mat1 := new(Matrix)
	mat1.AddRow([]float64{1, 2, 3, 4})
	mat1.AddRow([]float64{5, 6, 7, 8})
	mat1.AddRow([]float64{9, 10, 11, 12})
	mat1.AddRow([]float64{13, 14, 15, 16})

	vec1 := NewVector4(1, 2, 3, 4)

	matExpected := NewVector4(30, 70, 110, 150)
	matRes := mat1.Product(vec1)

	if !matRes.Equals(matExpected) {
		t.Errorf("Incorrect multiply vector, want: %+v, got: %+v", matExpected, matRes)
	}
}
