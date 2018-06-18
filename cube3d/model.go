package cube3d

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/nmaupu/gocube/compute"
	"github.com/nmaupu/gocube/data"
	"log"
	"math"
)

const (
	AngleOfView = 90
	Near        = .1
	Far         = 100.
	CamX        = 0.
	CamY        = 0.
	CamZ        = -20.
)

var (
	Mproj = GetProjectionMatrix(AngleOfView, Near, Far)
	Cam   = GetCameraTranslation(CamX, CamY, CamZ)
)

type cubie3d struct {
	Point     *compute.Matrix
	HexColor  string
	CubieSize float64
	CubeSize  int
	DirRight  *compute.Matrix
	DirDown   *compute.Matrix
}

func GetRotationMatrixX(rad float64) *compute.Matrix {
	m := new(compute.Matrix)
	m.AddRow([]float64{1, 0, 0, 0})
	m.AddRow([]float64{0, math.Cos(rad), -math.Sin(rad), 0})
	m.AddRow([]float64{0, math.Sin(rad), math.Cos(rad), 0})
	m.AddRow([]float64{0, 0, 0, 1})
	return m
}
func GetRotationMatrixY(rad float64) *compute.Matrix {
	m := new(compute.Matrix)
	m.AddRow([]float64{math.Cos(rad), 0, math.Sin(rad), 0})
	m.AddRow([]float64{0, 1, 0, 0})
	m.AddRow([]float64{-math.Sin(rad), 0, math.Cos(rad), 0})
	m.AddRow([]float64{0, 0, 0, 1})
	return m
}
func GetRotationMatrixZ(rad float64) *compute.Matrix {
	m := new(compute.Matrix)
	m.AddRow([]float64{math.Cos(rad), -math.Sin(rad), 0, 0})
	m.AddRow([]float64{math.Sin(rad), math.Cos(rad), 0, 0})
	m.AddRow([]float64{0, 0, 1, 0})
	m.AddRow([]float64{0, 0, 0, 1})
	return m
}

func GetTranslationMatrix(vec *compute.Matrix) *compute.Matrix {
	m := new(compute.Matrix)
	m.AddRow([]float64{1, 0, 0, vec.At(0, 0)})
	m.AddRow([]float64{0, 1, 0, vec.At(1, 0)})
	m.AddRow([]float64{0, 0, 1, vec.At(2, 0)})
	m.AddRow([]float64{0, 0, 0, 1})
	return m
}

func GetScaleMatrix(x, y, z float64) *compute.Matrix {
	m := new(compute.Matrix)
	m.AddRow([]float64{x, 0, 0, 0})
	m.AddRow([]float64{0, y, 0, 0})
	m.AddRow([]float64{0, 0, z, 0})
	m.AddRow([]float64{0, 0, 0, 1})
	return m
}

func getRad(deg float64) float64 {
	return (deg * math.Pi) / 180
}

func buildFace3d(cube *data.Cube, color string, radX, radY float64) []cubie3d {
	log.Printf("-- Building face %s", color)

	face := cube.Faces[color]
	ret := make([]cubie3d, 0)

	// In the real world coordinates, every points are going to be
	// between [-1,1] - left, right
	// and [-1,1] bottom, up
	// So, we are building a cube from bottom left piece at origin. Coordinates will be included in [0,2] for a 3x3
	// and at the end, we will center it such as its core will be at origin (0,0,0)
	// So for a 3X3, all coordinates will be in [-1.5,1.5]
	// We need to get that between [-1,1] and centered.
	// As a result, we will be translating the cube according to the number of cubies divided by 2 (dim/2)
	// And we will be scaling everything by 1/(dim/2) so all points will be included in [-1,1]

	// Note about 4x4 matrices
	// w=1 -> position in space
	// w=0 -> direction
	halfWidth := float64(cube.CubeSize) / 2. // i.e. For a 3x3, it's 1.5
	toOrigMat := compute.NewVector4(-halfWidth, -halfWidth, -halfWidth, 0)
	scale := 1. / halfWidth

	log.Println("toOrigMat: ", toOrigMat)
	log.Println("scale: ", scale)

	for i := 0; i < len(face.Colors); i++ {
		for j := 0; j < len(face.Colors[i]); j++ {
			c3d := cubie3d{
				HexColor:  face.Colors[i][j].HexColor,
				CubieSize: cube.CubieSize,
				CubeSize:  cube.CubeSize,
			}

			switch color {
			case "white":
				x := float64(j)
				y := float64(cube.CubeSize)
				z := -float64(cube.CubeSize) + float64(i)

				c3d.Point = compute.NewVector4(x, y, z, 1)
				c3d.DirRight = compute.NewVector4(1, 0, 0, 0)
				c3d.DirDown = compute.NewVector4(0, 0, 1.0, 0)

			case "green":
				x := float64(j)
				y := float64(cube.CubeSize) - float64(i)
				z := 0.

				c3d.Point = compute.NewVector4(x, y, z, 1)
				c3d.DirRight = compute.NewVector4(1, 0, 0, 0)
				c3d.DirDown = compute.NewVector4(0, -1, 0, 0)

			case "red":
				x := float64(cube.CubeSize)
				y := float64(cube.CubeSize) - float64(i)
				z := -float64(j)

				c3d.Point = compute.NewVector4(x, y, z, 1)
				c3d.DirRight = compute.NewVector4(0, 0, -1, 0)
				c3d.DirDown = compute.NewVector4(0, -1, 0, 0)
			}

			// Center the cube coordinates around origin
			c3d.Point = GetTranslationMatrix(toOrigMat).Product(c3d.Point)
			c3d.Point = GetScaleMatrix(scale, scale, scale).Product(c3d.Point)

			log.Printf("Resulting point (before rotations) = %+v", c3d.Point)

			c3d.Point = GetRotationMatrixX(radX).Product(c3d.Point)
			c3d.Point = GetRotationMatrixY(radY).Product(c3d.Point)

			c3d.DirRight = GetRotationMatrixX(radX).Product(c3d.DirRight).Normalize4()
			c3d.DirRight = GetRotationMatrixY(radY).Product(c3d.DirRight).Normalize4()

			c3d.DirDown = GetRotationMatrixX(radX).Product(c3d.DirDown).Normalize4()
			c3d.DirDown = GetRotationMatrixY(radY).Product(c3d.DirDown).Normalize4()

			ret = append(ret, c3d)
		}
	}

	return ret
}

// Convert a point to the original 3D coordinates plan
// to the drawing plan where Y is inverted
// Real world scene: X(right), Y(top), Z(back)
// Drawing plan / camera plan: X(right), Y(bottom), Z(front)
// Returns a 4D position vector (w=1), this is not really needed, it's done just by convention
func ConvertToDrawingPlan(vec *compute.Matrix, imgWidth, imgHeight int) *compute.Matrix {
	if vec.GetNbRows() != 4 && vec.GetNbCols() != 1 {
		panic("Parameter is not a 4D vector")
	}

	// Y and Z axis are inverted
	x := vec.At(0, 0)
	y := vec.At(1, 0)
	m := compute.NewVector4(
		(x+1)*.5*float64(imgWidth),
		(1-(y+1)*.5)*float64(imgHeight),
		0,
		1,
	)
	return m
}

func DrawCubie(ctx *gg.Context, x, y float64, c3d cubie3d) {
	// Need to scale the translation vector by the size of one cubie in the real world coordinates
	cubieScale := 1. / (float64(c3d.CubeSize) / 2.)

	m1 := c3d.Point
	m2 := GetTranslationMatrix(c3d.DirRight.ScalarMultiply(cubieScale)).Product(m1)
	m3 := GetTranslationMatrix(c3d.DirDown.ScalarMultiply(cubieScale)).Product(m2)
	m4 := GetTranslationMatrix(c3d.DirDown.ScalarMultiply(cubieScale)).Product(m1)

	// Real scale
	s := c3d.CubieSize
	scaleMat := GetScaleMatrix(s, s, s)

	// Scale and convert to drawing plan
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	m1 = ProjectPoint(scaleMat.Product(m1))
	m2 = ProjectPoint(scaleMat.Product(m2))
	m3 = ProjectPoint(scaleMat.Product(m3))
	m4 = ProjectPoint(scaleMat.Product(m4))
	p1 := ConvertToDrawingPlan(m1, ctx.Width(), ctx.Height())
	p2 := ConvertToDrawingPlan(m2, ctx.Width(), ctx.Height())
	p3 := ConvertToDrawingPlan(m3, ctx.Width(), ctx.Height())
	p4 := ConvertToDrawingPlan(m4, ctx.Width(), ctx.Height())

	ctx.SetLineWidth(6)
	ctx.MoveTo(x+p1.At(0, 0), y+p1.At(1, 0))
	ctx.LineTo(x+p2.At(0, 0), y+p2.At(1, 0))
	ctx.LineTo(x+p3.At(0, 0), y+p3.At(1, 0))
	ctx.LineTo(x+p4.At(0, 0), y+p4.At(1, 0))
	ctx.LineTo(x+p1.At(0, 0), y+p1.At(1, 0))
	ctx.SetHexColor(c3d.HexColor)
	ctx.FillPreserve()
	ctx.SetHexColor("#000000")
	ctx.Stroke()
}

func DrawFace(ctx *gg.Context, x, y float64, c3ds []cubie3d) {
	for _, c3d := range c3ds {
		DrawCubie(ctx, x, y, c3d)
	}
}

func DrawAxes(ctx *gg.Context, x, y, width, radX, radY float64) *gg.Context {
	// Origins' points
	poX := compute.NewVector4(0, 0, 0, 1)
	poY := compute.NewVector4(0, 0, 0, 1)
	poZ := compute.NewVector4(0, 0, 0, 1)

	// Real world coords axes
	arrowScale := .045
	rwAxisX := GetTranslationMatrix(compute.NewVector3(1, 0, 0)).Product(poX)
	rwArrowX1 := GetTranslationMatrix(compute.NewVector3(-arrowScale, 0, arrowScale)).Product(rwAxisX)
	rwArrowX2 := GetTranslationMatrix(compute.NewVector3(-arrowScale, 0, -arrowScale)).Product(rwAxisX)

	rwAxisY := GetTranslationMatrix(compute.NewVector3(0, 1, 0)).Product(poY)
	rwArrowY1 := GetTranslationMatrix(compute.NewVector3(-arrowScale, -arrowScale, 0)).Product(rwAxisY)
	rwArrowY2 := GetTranslationMatrix(compute.NewVector3(arrowScale, -arrowScale, 0)).Product(rwAxisY)

	rwAxisZ := GetTranslationMatrix(compute.NewVector3(0, 0, -1)).Product(poZ)
	rwArrowZ1 := GetTranslationMatrix(compute.NewVector3(-arrowScale, 0, arrowScale)).Product(rwAxisZ)
	rwArrowZ2 := GetTranslationMatrix(compute.NewVector3(arrowScale, 0, arrowScale)).Product(rwAxisZ)

	rwAxisX = GetRotationMatrixX(radX).Product(rwAxisX)
	rwAxisX = GetRotationMatrixY(radY).Product(rwAxisX)
	rwArrowX1 = GetRotationMatrixX(radX).Product(rwArrowX1)
	rwArrowX1 = GetRotationMatrixY(radY).Product(rwArrowX1)
	rwArrowX2 = GetRotationMatrixX(radX).Product(rwArrowX2)
	rwArrowX2 = GetRotationMatrixY(radY).Product(rwArrowX2)

	rwAxisY = GetRotationMatrixX(radX).Product(rwAxisY)
	rwAxisY = GetRotationMatrixY(radY).Product(rwAxisY)
	rwArrowY1 = GetRotationMatrixX(radX).Product(rwArrowY1)
	rwArrowY1 = GetRotationMatrixY(radY).Product(rwArrowY1)
	rwArrowY2 = GetRotationMatrixX(radX).Product(rwArrowY2)
	rwArrowY2 = GetRotationMatrixY(radY).Product(rwArrowY2)

	rwAxisZ = GetRotationMatrixX(radX).Product(rwAxisZ)
	rwAxisZ = GetRotationMatrixY(radY).Product(rwAxisZ)
	rwArrowZ1 = GetRotationMatrixX(radX).Product(rwArrowZ1)
	rwArrowZ1 = GetRotationMatrixY(radY).Product(rwArrowZ1)
	rwArrowZ2 = GetRotationMatrixX(radX).Product(rwArrowZ2)
	rwArrowZ2 = GetRotationMatrixY(radY).Product(rwArrowZ2)

	// Axis
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	scaleMat := GetScaleMatrix(width, width, width)
	aX := ConvertToDrawingPlan(ProjectPoint(scaleMat.Product(rwAxisX)), ctx.Width(), ctx.Height())
	arrX1 := ConvertToDrawingPlan(ProjectPoint(scaleMat.Product(rwArrowX1)), ctx.Width(), ctx.Height())
	arrX2 := ConvertToDrawingPlan(ProjectPoint(scaleMat.Product(rwArrowX2)), ctx.Width(), ctx.Height())

	aY := ConvertToDrawingPlan(ProjectPoint(scaleMat.Product(rwAxisY)), ctx.Width(), ctx.Height())
	arrY1 := ConvertToDrawingPlan(ProjectPoint(scaleMat.Product(rwArrowY1)), ctx.Width(), ctx.Height())
	arrY2 := ConvertToDrawingPlan(ProjectPoint(scaleMat.Product(rwArrowY2)), ctx.Width(), ctx.Height())

	aZ := ConvertToDrawingPlan(ProjectPoint(scaleMat.Product(rwAxisZ)), ctx.Width(), ctx.Height())
	arrZ1 := ConvertToDrawingPlan(ProjectPoint(scaleMat.Product(rwArrowZ1)), ctx.Width(), ctx.Height())
	arrZ2 := ConvertToDrawingPlan(ProjectPoint(scaleMat.Product(rwArrowZ2)), ctx.Width(), ctx.Height())

	origX := ConvertToDrawingPlan(ProjectPoint(poX), ctx.Width(), ctx.Height())
	origY := ConvertToDrawingPlan(ProjectPoint(poY), ctx.Width(), ctx.Height())
	origZ := ConvertToDrawingPlan(ProjectPoint(poZ), ctx.Width(), ctx.Height())

	// Draw axes
	ctx.SetLineWidth(5)

	ctx.SetHexColor("#FF0000")
	ctx.DrawLine(origX.At(0, 0), origX.At(1, 0), aX.At(0, 0), aX.At(1, 0))
	ctx.DrawLine(aX.At(0, 0), aX.At(1, 0), arrX1.At(0, 0), arrX1.At(1, 0))
	ctx.DrawLine(aX.At(0, 0), aX.At(1, 0), arrX2.At(0, 0), arrX2.At(1, 0))
	ctx.Stroke()

	ctx.SetHexColor("#00FF00")
	ctx.DrawLine(origY.At(0, 0), origY.At(1, 0), aY.At(0, 0), aY.At(1, 0))
	ctx.DrawLine(aY.At(0, 0), aY.At(1, 0), arrY1.At(0, 0), arrY1.At(1, 0))
	ctx.DrawLine(aY.At(0, 0), aY.At(1, 0), arrY2.At(0, 0), arrY2.At(1, 0))
	ctx.Stroke()

	ctx.SetHexColor("#0000FF")
	ctx.DrawLine(origZ.At(0, 0), origZ.At(0, 0), aZ.At(0, 0), aZ.At(1, 0))
	ctx.DrawLine(aZ.At(0, 0), aZ.At(1, 0), arrZ1.At(0, 0), arrZ1.At(1, 0))
	ctx.DrawLine(aZ.At(0, 0), aZ.At(1, 0), arrZ2.At(0, 0), arrZ2.At(1, 0))
	ctx.Stroke()

	return ctx
}

// Perspective projection matrix given by:
// angleOfView = FOV in degrees
// n = near clipping plane
// f = far clipping plane
func GetProjectionMatrix(angleOfView, n, f float64) *compute.Matrix {
	ret := new(compute.Matrix)
	radFov := getRad(angleOfView)
	aspect := 1.

	scaleTan := math.Tan(radFov * .5)
	ret.AddRow([]float64{1 / (aspect * scaleTan), 0, 0, 0})
	ret.AddRow([]float64{0, 1 / scaleTan, 0, 0})
	ret.AddRow([]float64{0, 0, -((f + n) / (f - n)), -2 * f * n / (f - n)})
	ret.AddRow([]float64{0, 0, -1, 0})

	return ret
}

// Multiply a 4D vector by a 4x4 matrix
// and divide by w (perspective occurs here) and w != 1
// returns a vector 4D
func MultPointMatrix(m *compute.Matrix, v *compute.Matrix) *compute.Matrix {
	ret := m.Product(v)

	w := ret.At(3, 0)
	// normalize if w is different than 1 (convert from homogeneous to Cartesian coordinates)
	if w != 1 {
		ret.Data[0][0] /= w
		ret.Data[1][0] /= w
		ret.Data[2][0] /= w
	}

	return ret
}

// Get translation matrix for camera "the eye"
func GetCameraTranslation(x, y, z float64) *compute.Matrix {
	return GetTranslationMatrix(compute.NewVector3(x, y, z))
}

// Project a point on the plan using perspective projection
// p, the point must be a 4D vector
// Returns a 4D vector point. Only x and y are useful ;-)
func ProjectPoint(p *compute.Matrix) *compute.Matrix {
	// ret = cam x Mproj x p
	ret := MultPointMatrix(Cam, p)
	ret = MultPointMatrix(Mproj, ret)

	log.Printf("Projecting point is : %+v", ret)

	return ret

	// x and y must be in [-1,1] to be rendered
	x := ret.At(0, 0)
	y := ret.At(1, 0)
	if x >= -1 && x <= 1 && y >= -1 && y <= 1 {
		return ret
	} else {
		panic(fmt.Sprintf("Point cannot be rendered point=%+v", ret))
	}
}

func DrawCube(ctx *gg.Context, x, y float64, cube *data.Cube) *gg.Context {
	var face3dMatrices []cubie3d

	radX := getRad(35.264)
	radY := -getRad(45)
	radX = 0
	radY = 0

	face3dMatrices = buildFace3d(cube, "white", radX, radY)
	DrawFace(ctx, x, y, face3dMatrices)
	face3dMatrices = buildFace3d(cube, "red", radX, radY)
	DrawFace(ctx, x, y, face3dMatrices)
	face3dMatrices = buildFace3d(cube, "green", radX, radY)
	DrawFace(ctx, x, y, face3dMatrices)

	DrawAxes(ctx, x, y, 10, radX, radY)

	return ctx
}
