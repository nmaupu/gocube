package data

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/fogleman/gg"
)

// Faces are corresponding to the following diagram
//       W
//    O  G  R  B
//       Y
// using colors to identify the face

type Face struct {
	Colors [][]Color
}

func NewFace(cubeSize int, color Color) *Face {
	f := Face{}
	f.Colors = make([][]Color, cubeSize)

	for i := 0; i < cubeSize; i++ {
		row := make([]Color, cubeSize)
		f.Colors[i] = row

		for j := 0; j < cubeSize; j++ {
			colorCopy := color
			format := "%s%d"
			if cubeSize*cubeSize > 9 {
				format = "%s%02d"
			}
			colorCopy.Debug = fmt.Sprintf(format, color.ShortName, i*cubeSize+j+1)
			row[j] = colorCopy
		}
	}

	return &f
}

func (f Face) GetCenterColor() (*Color, error) {
	len := len(f.Colors)
	if len%2 == 0 {
		return nil, errors.New(fmt.Sprintf("Cannot get center piece on a %dx%d cube !", len, len))
	}

	return &(f.Colors[len/2][len/2]), nil
}

func (f Face) String() string {
	var buf bytes.Buffer
	cubeSize := len(f.Colors)

	for i := 0; i < cubeSize; i++ {
		buf.WriteString(f.StringRow(i))
		buf.WriteString("\n")
	}

	return buf.String()
}

func (f Face) StringRow(n int) string {
	var buf bytes.Buffer
	cubeSize := len(f.Colors[n])

	for i := 0; i < cubeSize; i++ {
		buf.WriteString(f.Colors[n][i].String())
		if i < cubeSize-1 {
			buf.WriteString(" ")
		} else {
			buf.WriteString("|")
		}
	}

	return buf.String()
}

func (f Face) Draw(ctx *gg.Context, x, y, cubieSize float64) {
	cubeSize := len(f.Colors)
	for i := 0; i < cubeSize; i++ {
		for j := 0; j < cubeSize; j++ {
			f.Colors[i][j].Draw(
				ctx,
				x+float64(j)*cubieSize,
				y+float64(i)*cubieSize,
				cubieSize)
		}
	}
}

// Draw one row only
func (f Face) DrawRow(ctx *gg.Context, row int, x, y, cubieSize float64) {
	cubeSize := len(f.Colors)
	for j := 0; j < cubeSize; j++ {
		f.Colors[row][j].DrawHalfH(
			ctx,
			x+float64(j)*cubieSize,
			y,
			cubieSize)
	}
}

// Draw one row only in column
func (f Face) DrawRowCol(ctx *gg.Context, row int, x, y, cubieSize float64) {
	cubeSize := len(f.Colors)
	for j := 0; j < cubeSize; j++ {
		f.Colors[row][j].DrawHalfV(
			ctx,
			x,
			y+float64(j)*cubieSize,
			cubieSize)
	}
}

func (f Face) Copy() *Face {
	cubeSize := len(f.Colors)
	ret := NewFace(cubeSize, Colors["white"])

	for i := 0; i < cubeSize; i++ {
		for j := 0; j < cubeSize; j++ {
			ret.Colors[i][j] = f.Colors[i][j]
		}
	}

	return ret
}

// Flip a face (useful to display blue face in a 2D representation)
func (f Face) FlipVertical() Face {
	cubeSize := len(f.Colors)
	orig := f.Copy()
	for i := 0; i < cubeSize; i++ {
		for j := 0; j < cubeSize; j++ {
			f.Colors[i][j] = orig.Colors[i][cubeSize-j-1]
		}
	}

	return f
}

func (f Face) FlipHorizontal() Face {
	cubeSize := len(f.Colors)
	orig := f.Copy()
	for i := 0; i < cubeSize; i++ {
		for j := 0; j < cubeSize; j++ {
			f.Colors[i][j] = orig.Colors[cubeSize-i-1][j]
		}
	}

	return f
}

func (f Face) Flip() Face {
	return f.FlipHorizontal().FlipVertical()
}

// Replace current given col with the one from the given face
func (f Face) ReplaceCol(f2 Face, col int) Face {
	cubeSize := len(f.Colors)
	for i := 0; i < cubeSize; i++ {
		f.Colors[i][col] = f2.Colors[i][col]
	}

	return f
}

func (f Face) ReplaceRow(f2 Face, row int) Face {
	cubeSize := len(f.Colors)
	for c := 0; c < cubeSize; c++ {
		f.Colors[row][c] = f2.Colors[row][c]
	}

	return f
}

func (f Face) RotateClockwise() Face {
	cubeSize := len(f.Colors)
	copy := f.Copy()

	for c := 0; c < cubeSize; c++ {
		for i := cubeSize - 1; i >= 0; i-- {
			f.Colors[c][cubeSize-1-i] = copy.Colors[i][c]
		}
	}

	return f
}

func (f Face) RotateAntiClockwise() Face {
	// Make 3 times RotateClockwise
	for i := 0; i < 3; i++ {
		f.RotateClockwise()
	}

	return f
}
