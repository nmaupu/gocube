package math3D

import (
	"math"
)

func NewVector3(x, y, z float64) *Matrix {
	ret := new(Matrix)
	ret.AddRow([]float64{x, y, z})
	return ret.Transpose()
}

func NewVector4(x, y, z, w float64) *Matrix {
	ret := new(Matrix)
	ret.AddRow([]float64{x, y, z, w})
	return ret.Transpose()
}

func (m *Matrix) Normalize4() *Matrix {
	length := m.VectorLength()
	return NewVector4(
		m.At(0, 0)/length,
		m.At(1, 0)/length,
		m.At(2, 0)/length,
		0,
	)
}

func (m *Matrix) VectorLength() float64 {
	adds := 0.

	// For a vector4, w is set to 0 so
	// no prob adding 0*0
	for i := 0; i < m.GetNbRows(); i++ {
		adds += m.At(i, 0) * m.At(i, 0)
	}

	return math.Sqrt(adds)
}
