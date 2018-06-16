package cube3d

import (
	"github.com/fogleman/gg"
	"github.com/nmaupu/gocube/compute"
	"github.com/nmaupu/gocube/data"
	"math"
)

type cubie3d struct {
	Point     *compute.Matrix
	HexColor  string
	CubieSize float64
	DirRight  *compute.Matrix
	DirDown   *compute.Matrix
}

func GetRotationMatrixX(rad float64) *compute.Matrix {
	m := new(compute.Matrix)
	m.AddRow([]float64{1, 0, 0})
	m.AddRow([]float64{0, math.Cos(rad), -1 * math.Sin(rad)})
	m.AddRow([]float64{0, math.Sin(rad), math.Cos(rad)})
	return m
}
func GetRotationMatrixY(rad float64) *compute.Matrix {
	m := new(compute.Matrix)
	m.AddRow([]float64{math.Cos(rad), 0, math.Sin(rad)})
	m.AddRow([]float64{0, 1, 0})
	m.AddRow([]float64{-1 * math.Sin(rad), 0, math.Cos(rad)})
	return m
}
func GetRotationMatrixZ(rad float64) *compute.Matrix {
	m := new(compute.Matrix)
	m.AddRow([]float64{math.Cos(rad), -1 * math.Sin(rad), 0})
	m.AddRow([]float64{math.Sin(rad), math.Cos(rad), 0})
	m.AddRow([]float64{0, 0, 1})
	return m
}

func GetOrthographicProjectionMatrix() *compute.Matrix {
	m := new(compute.Matrix)
	m.AddRow([]float64{1, 0, 0})
	m.AddRow([]float64{0, 1, 0})
	m.AddRow([]float64{0, 0, 0})
	return m
}

func getRad(deg float64) float64 {
	return (deg * math.Pi) / 180
}

func buildFace3d(cube *data.Cube, color string) []cubie3d {
	ret := make([]cubie3d, 0)
	face := cube.Faces[color]
	if color == "orange" {
		face = face.FlipVertical()
	}
	if color == "yellow" {
		face = face.FlipHorizontal()
	}

	for i := 0; i < len(face.Colors); i++ {
		for j := 0; j < len(face.Colors[i]); j++ {
			c3d := cubie3d{
				HexColor:  face.Colors[i][j].HexColor,
				CubieSize: cube.CubieSize,
			}

			switch color {
			case "white":
				x := float64(j) * cube.CubieSize
				y := 0.0
				z := float64(i) * cube.CubieSize

				translationZ := compute.NewMatrix([]float64{0, 0, -1 * float64(cube.CubeSize-1) * cube.CubieSize})
				c3d.Point = compute.NewMatrix([]float64{x, y, z}).Add(translationZ)
				c3d.DirRight = compute.NewMatrix([]float64{float64(cube.CubieSize), 0, 0})
				c3d.DirDown = compute.NewMatrix([]float64{0, 0, -1.0 * float64(cube.CubieSize)})

			case "yellow":
				x := float64(j) * cube.CubieSize
				y := float64(cube.CubeSize) * cube.CubieSize
				z := float64(i) * cube.CubieSize

				translationZ := compute.NewMatrix([]float64{0, 0, -1 * float64(cube.CubeSize-1) * cube.CubieSize})
				c3d.Point = compute.NewMatrix([]float64{x, y, z}).Add(translationZ)
				c3d.DirRight = compute.NewMatrix([]float64{float64(cube.CubieSize), 0, 0})
				c3d.DirDown = compute.NewMatrix([]float64{0, 0, -1.0 * float64(cube.CubieSize)})

			case "green":
				x := float64(j) * cube.CubieSize
				y := float64(i) * cube.CubieSize
				z := 0.0

				c3d.Point = compute.NewMatrix([]float64{x, y, z})
				c3d.DirRight = compute.NewMatrix([]float64{float64(cube.CubieSize), 0, 0})
				c3d.DirDown = compute.NewMatrix([]float64{0, 1 * float64(cube.CubieSize), 0})

			case "red":
				x := float64(cube.CubeSize) * cube.CubieSize
				y := float64(i) * cube.CubieSize
				z := -1 * float64(j) * cube.CubieSize

				c3d.Point = compute.NewMatrix([]float64{x, y, z})
				c3d.DirRight = compute.NewMatrix([]float64{0, 0, -1 * float64(cube.CubieSize)})
				c3d.DirDown = compute.NewMatrix([]float64{0, float64(cube.CubieSize), 0})

			case "orange":
				x := 0.0
				y := float64(i) * cube.CubieSize
				z := -1 * float64(j) * cube.CubieSize

				c3d.Point = compute.NewMatrix([]float64{x, y, z})
				c3d.DirRight = compute.NewMatrix([]float64{0, 0, -1 * float64(cube.CubieSize)})
				c3d.DirDown = compute.NewMatrix([]float64{0, float64(cube.CubieSize), 0})

			}

			ret = append(ret, c3d)
		}
	}

	return ret
}

func DrawCubie(ctx *gg.Context, px, py float64, c3d cubie3d, radX, radY float64) {
	m1 := c3d.Point
	m2 := c3d.Point.Add(c3d.DirRight)
	m3 := m1.Add(c3d.DirDown)
	m4 := m2.Add(c3d.DirDown)

	rot1 := m1.
		Product(GetRotationMatrixY(radY)).
		Product(GetRotationMatrixX(radX)).
		Product(GetOrthographicProjectionMatrix())

	rot2 := m2.
		Product(GetRotationMatrixY(radY)).
		Product(GetRotationMatrixX(radX)).
		Product(GetOrthographicProjectionMatrix())

	rot3 := m3.
		Product(GetRotationMatrixY(radY)).
		Product(GetRotationMatrixX(radX)).
		Product(GetOrthographicProjectionMatrix())

	rot4 := m4.
		Product(GetRotationMatrixY(radY)).
		Product(GetRotationMatrixX(radX)).
		Product(GetOrthographicProjectionMatrix())

	// Trace cubie using lines
	ctx.SetLineWidth(c3d.CubieSize * 5 / 100)
	ctx.MoveTo(px+rot1.At(0, 0), py+rot1.At(0, 1))
	ctx.LineTo(px+rot2.At(0, 0), py+rot2.At(0, 1))
	ctx.LineTo(px+rot4.At(0, 0), py+rot4.At(0, 1))
	ctx.LineTo(px+rot3.At(0, 0), py+rot3.At(0, 1))
	ctx.LineTo(px+rot1.At(0, 0), py+rot1.At(0, 1))
	ctx.SetHexColor(c3d.HexColor)
	ctx.FillPreserve()
	ctx.SetHexColor("#000000")
	ctx.Stroke()
}

func DrawFace(ctx *gg.Context, x, y float64, c3ds []cubie3d, radX, radY float64) {
	for _, c3d := range c3ds {
		DrawCubie(ctx, x, y, c3d, radX, radY)
	}
}

func DrawCube(ctx *gg.Context, x, y float64, cube *data.Cube) *gg.Context {
	var face3dMatrices []cubie3d

	// Draw axes
	//ctx.NewSubPath()
	//ctx.DrawLine(x, y, x+400, y)
	//ctx.DrawLine(x, y, x, y+400)
	//ctx.Stroke()

	face3dMatrices = buildFace3d(cube, "white")
	DrawFace(ctx, x, y, face3dMatrices, getRad(32), getRad(45))
	face3dMatrices = buildFace3d(cube, "green")
	DrawFace(ctx, x, y, face3dMatrices, getRad(32), getRad(45))
	face3dMatrices = buildFace3d(cube, "red")
	DrawFace(ctx, x, y, face3dMatrices, getRad(32), getRad(45))

	return ctx
}
