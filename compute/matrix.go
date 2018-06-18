package compute

import (
	"fmt"
	"math"
)

type Matrix struct {
	Data [][]float64
}

func NewMatrix(row []float64) *Matrix {
	ret := new(Matrix)
	ret.AddRow(row)
	return ret
}

func NewVector(coords []float64) *Matrix {
	ret := new(Matrix)
	ret.AddRow(coords)
	return ret.Transpose()
}

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

func (m *Matrix) AddRow(row []float64) *Matrix {
	m.Data = append(m.Data, row)
	return m
}

func (m *Matrix) At(row, col int) float64 {
	return m.Data[row][col]
}

func (m *Matrix) GetNbRows() int {
	return len(m.Data)
}

func (m *Matrix) GetNbCols() int {
	return len(m.Data[0])
}

func (m *Matrix) GetRow(i int) *Matrix {
	ret := new(Matrix)
	ret.AddRow(m.Data[i])
	return ret
}

func (m *Matrix) GetCol(j int) *Matrix {
	ret := new(Matrix)
	col := make([]float64, 0)
	for i := 0; i < m.GetNbRows(); i++ {
		col = append(col, m.Data[i][j])
	}
	ret.AddRow(col)
	return ret
}

func (m *Matrix) Copy() *Matrix {
	ret := Matrix{}
	copy(m.Data, ret.Data)
	return &ret
}

func (m *Matrix) Equals(m2 *Matrix) bool {
	if len(m.Data) != len(m2.Data) {
		return false
	}

	for i := 0; i < len(m.Data); i++ {
		for j := 0; j < len(m.Data[i]); j++ {
			if m.Data[i][j] != m2.Data[i][j] {
				return false
			}
		}
	}

	return true
}

func (m *Matrix) Transpose() *Matrix {
	ret := new(Matrix)

	for j := 0; j < m.GetNbCols(); j++ {
		col := make([]float64, 0)
		for i := 0; i < m.GetNbRows(); i++ {
			col = append(col, m.At(i, j))
		}
		ret.AddRow(col)
	}

	return ret
}

func (m1 *Matrix) Product(m2 *Matrix) *Matrix {
	if m1.GetNbCols() != m2.GetNbRows() {
		panic(fmt.Sprintf("Unable to multiply matrices : %+v, %+v", m1, m2))
	}

	matProd := new(Matrix)

	for i := 0; i < m1.GetNbRows(); i++ {
		curRow := m1.GetRow(i)
		rowProd := make([]float64, 0)

		for j := 0; j < m2.GetNbCols(); j++ {
			curCol := m2.GetCol(j)
			val := mult(curRow.Data[0], curCol.Data[0])
			rowProd = append(rowProd, val)
		}

		matProd.AddRow(rowProd)
	}

	return matProd
}

func (m *Matrix) ScalarMultiply(s float64) *Matrix {
	ret := new(Matrix)

	for i := 0; i < m.GetNbRows(); i++ {
		row := make([]float64, 0)

		for j := 0; j < m.GetNbCols(); j++ {
			row = append(row, s*m.Data[i][j])
		}

		ret.AddRow(row)
	}

	return ret
}

func mult(a1, a2 []float64) float64 {
	ret := 0.0

	for i := 0; i < len(a1); i++ {
		ret += a1[i] * a2[i]
	}

	return ret
}

func (m1 *Matrix) Add(m2 *Matrix) *Matrix {
	if m1.GetNbCols() != m2.GetNbCols() || m1.GetNbRows() != m2.GetNbRows() {
		panic(fmt.Sprintf("Unable to add matrices: %+v, %+v", m1, m2))
	}

	m := new(Matrix)
	for i := 0; i < len(m1.Data); i++ {
		row := make([]float64, 0)
		for j := 0; j < len(m1.Data[i]); j++ {
			row = append(row, m1.Data[i][j]+m2.Data[i][j])
		}
		m.AddRow(row)
	}

	return m
}
