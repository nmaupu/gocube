package cube3d

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/nmaupu/gocube/data"
	"github.com/oelmekki/matrix"
	"math"
)

func BuildFace3d(face data.Face, cubieSize float64) []matrix.Matrix {
	var mat []matrix.Matrix

	for i := 0; i < len(face.Colors); i++ {
		for j := 0; j < len(face.Colors[i]); j++ {
			m, _ := matrix.Build(
				matrix.Builder{
					matrix.Row{float64(i) * cubieSize, float64(j) * cubieSize, 0},
				},
			)

			mat = append(mat, m)
		}
	}

	return mat
}

func GetRotationMatrixX(rad float64) matrix.Matrix {
	m, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{1, 0, 0},
			matrix.Row{0, math.Cos(rad), -1 * math.Sin(rad)},
			matrix.Row{0, math.Sin(rad), math.Cos(rad)},
		},
	)
	return m
}
func GetRotationMatrixY(rad float64) matrix.Matrix {
	m, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{math.Cos(rad), 0, math.Sin(rad)},
			matrix.Row{0, 1, 0},
			matrix.Row{-1 * math.Sin(rad), 0, math.Cos(rad)},
		},
	)
	return m
}
func GetRotationMatrixZ(rad float64) matrix.Matrix {
	m, _ := matrix.Build(
		matrix.Builder{
			matrix.Row{math.Cos(rad), -1 * math.Sin(rad), 0},
			matrix.Row{math.Sin(rad), math.Cos(rad), 0},
			matrix.Row{0, 0, 1},
		},
	)
	return m
}

func DrawCubie(ctx *gg.Context, px, py float64, mat matrix.Matrix, cubieSize float64) {
	x := mat.At(0, 0)
	y := mat.At(0, 1)
	z := mat.At(0, 2)
	m1 := mat
	m2, _ := matrix.Build(matrix.Builder{
		matrix.Row{x + cubieSize, y, z},
	})
	m3, _ := matrix.Build(matrix.Builder{
		matrix.Row{x, y + cubieSize, z},
	})
	m4, _ := matrix.Build(matrix.Builder{
		matrix.Row{x + cubieSize, y + cubieSize, z},
	})

	rad := 45 * math.Pi / 180
	rot1, _ := m1.DotProduct(GetRotationMatrixY(rad))
	rot1, _ = rot1.DotProduct(GetRotationMatrixX(32 * math.Pi / 180))
	rot2, _ := m2.DotProduct(GetRotationMatrixY(rad))
	rot2, _ = rot2.DotProduct(GetRotationMatrixX(32 * math.Pi / 180))
	rot3, _ := m3.DotProduct(GetRotationMatrixY(rad))
	rot3, _ = rot3.DotProduct(GetRotationMatrixX(32 * math.Pi / 180))
	rot4, _ := m4.DotProduct(GetRotationMatrixY(rad))
	rot4, _ = rot4.DotProduct(GetRotationMatrixX(32 * math.Pi / 180))

	fmt.Println("All resulting points:")
	fmt.Println(rot1)
	fmt.Println(rot2)
	fmt.Println(rot3)
	fmt.Println(rot4)

	// Trace cubie using lines
	ctx.DrawLine(px+rot1.At(0, 0), py+rot1.At(0, 1), px+rot2.At(0, 0), py+rot2.At(0, 1))
	ctx.DrawLine(px+rot2.At(0, 0), py+rot2.At(0, 1), px+rot4.At(0, 0), py+rot4.At(0, 1))
	ctx.DrawLine(px+rot4.At(0, 0), py+rot4.At(0, 1), px+rot3.At(0, 0), py+rot3.At(0, 1))
	ctx.DrawLine(px+rot3.At(0, 0), py+rot3.At(0, 1), px+rot1.At(0, 0), py+rot1.At(0, 1))
	ctx.Stroke()
}

func DrawFace(ctx *gg.Context, x, y float64, face []matrix.Matrix, cubieSize float64) {
	for _, f := range face {
		DrawCubie(ctx, x, y, f, cubieSize)
	}
}
