package data

import (
	"bytes"
	"strings"
)

type Cube struct {
	Faces    map[string]Face
	CubeSize int
}

func NewCube(cubeSize int) *Cube {
	c := Cube{
		Faces:    make(map[string]Face),
		CubeSize: cubeSize,
	}

	for k, color := range Colors {
		c.Faces[k] = *NewFace(cubeSize, color)
	}

	return &c
}

func (c Cube) String() string {
	// Representation as string is as follow:
	//       w w w
	//       w w w
	//       w w w
	// o o o g g g r r r b b b
	// o o o g g g r r r b b b
	// o o o g g g r r r b b b
	//       y y y
	//       y y y
	//       y y y

	var buf bytes.Buffer
	filler := strings.Repeat(
		" ", len(c.Faces["white"].StringRow(0))-1)
	filler += "|"

	for i := 0; i < c.CubeSize; i++ {
		buf.WriteString(filler)
		buf.WriteString(c.Faces["white"].StringRow(i))
		buf.WriteString("\n")
	}
	for i := 0; i < c.CubeSize; i++ {
		buf.WriteString(c.Faces["orange"].StringRow(i))
		buf.WriteString(c.Faces["green"].StringRow(i))
		buf.WriteString(c.Faces["red"].StringRow(i))
		buf.WriteString(c.Faces["blue"].StringRow(i))
		buf.WriteString("\n")
	}
	for i := 0; i < c.CubeSize; i++ {
		buf.WriteString(filler)
		buf.WriteString(c.Faces["yellow"].StringRow(i))
		buf.WriteString("\n")
	}

	return buf.String()
}

func (c *Cube) R() *Cube {
	colIndex := c.CubeSize - 1
	whiteCopy := *(c.Faces["white"].Copy())
	c.Faces["blue"].Flip()
	c.Faces["white"].ReplaceCol(c.Faces["green"], colIndex)
	c.Faces["green"].ReplaceCol(c.Faces["yellow"], colIndex)
	c.Faces["yellow"].ReplaceCol(c.Faces["blue"], colIndex)
	c.Faces["blue"].ReplaceCol(whiteCopy, colIndex)
	c.Faces["red"].RotateClockwise()
	c.Faces["blue"].Flip()

	return c
}

func (c *Cube) R2() *Cube {
	return c.R().R()
}

func (c *Cube) Rp() *Cube {
	colIndex := c.CubeSize - 1
	yellowCopy := *(c.Faces["yellow"].Copy())
	c.Faces["blue"].Flip()
	c.Faces["yellow"].ReplaceCol(c.Faces["green"], colIndex)
	c.Faces["green"].ReplaceCol(c.Faces["white"], colIndex)
	c.Faces["white"].ReplaceCol(c.Faces["blue"], colIndex)
	c.Faces["blue"].ReplaceCol(yellowCopy, colIndex)
	c.Faces["red"].RotateAntiClockwise()
	c.Faces["blue"].Flip()

	return c
}

func (c *Cube) L() *Cube {
	yellowCopy := *(c.Faces["yellow"].Copy())
	c.Faces["blue"].Flip()
	c.Faces["yellow"].ReplaceCol(c.Faces["green"], 0)
	c.Faces["green"].ReplaceCol(c.Faces["white"], 0)
	c.Faces["white"].ReplaceCol(c.Faces["blue"], 0)
	c.Faces["blue"].ReplaceCol(yellowCopy, 0)
	c.Faces["orange"].RotateClockwise()
	c.Faces["blue"].Flip()

	return c
}

func (c *Cube) L2() *Cube {
	return c.L().L()
}

func (c *Cube) Lp() *Cube {
	whiteCopy := *(c.Faces["white"].Copy())
	c.Faces["blue"].Flip()
	c.Faces["white"].ReplaceCol(c.Faces["green"], 0)
	c.Faces["green"].ReplaceCol(c.Faces["yellow"], 0)
	c.Faces["yellow"].ReplaceCol(c.Faces["blue"], 0)
	c.Faces["blue"].ReplaceCol(whiteCopy, 0)
	c.Faces["orange"].RotateAntiClockwise()
	c.Faces["blue"].Flip()

	return c
}

func (c *Cube) D() *Cube {
	rowIndex := c.CubeSize - 1
	greenCopy := *(c.Faces["green"].Copy())
	c.Faces["green"].ReplaceRow(c.Faces["orange"], rowIndex)
	c.Faces["orange"].ReplaceRow(c.Faces["blue"], rowIndex)
	c.Faces["blue"].ReplaceRow(c.Faces["red"], rowIndex)
	c.Faces["red"].ReplaceRow(greenCopy, rowIndex)
	c.Faces["yellow"].RotateClockwise()

	return c
}

func (c *Cube) D2() *Cube {
	return c.D().D()
}

func (c *Cube) Dp() *Cube {
	rowIndex := c.CubeSize - 1
	greenCopy := *(c.Faces["green"].Copy())
	c.Faces["green"].ReplaceRow(c.Faces["red"], rowIndex)
	c.Faces["red"].ReplaceRow(c.Faces["blue"], rowIndex)
	c.Faces["blue"].ReplaceRow(c.Faces["orange"], rowIndex)
	c.Faces["orange"].ReplaceRow(greenCopy, rowIndex)
	c.Faces["yellow"].RotateAntiClockwise()

	return c
}

func (c *Cube) U() *Cube {
	greenCopy := *(c.Faces["green"].Copy())
	c.Faces["green"].ReplaceRow(c.Faces["red"], 0)
	c.Faces["red"].ReplaceRow(c.Faces["blue"], 0)
	c.Faces["blue"].ReplaceRow(c.Faces["orange"], 0)
	c.Faces["orange"].ReplaceRow(greenCopy, 0)
	c.Faces["white"].RotateClockwise()

	return c
}

func (c *Cube) U2() *Cube {
	return c.U().U()
}

func (c *Cube) Up() *Cube {
	greenCopy := *(c.Faces["green"].Copy())
	c.Faces["green"].ReplaceRow(c.Faces["orange"], 0)
	c.Faces["orange"].ReplaceRow(c.Faces["blue"], 0)
	c.Faces["blue"].ReplaceRow(c.Faces["red"], 0)
	c.Faces["red"].ReplaceRow(greenCopy, 0)
	c.Faces["white"].RotateAntiClockwise()

	return c
}

func (c *Cube) F() *Cube {
	index := c.CubeSize - 1
	whiteCopy := *(c.Faces["white"].Copy())
	orangeCopy := *(c.Faces["orange"].Copy())
	yellowCopy := *(c.Faces["yellow"].Copy())
	redCopy := *(c.Faces["red"].Copy())
	c.Faces["white"].ReplaceRow(orangeCopy.RotateClockwise(), index)
	c.Faces["orange"].ReplaceCol(yellowCopy.RotateClockwise(), index)
	c.Faces["yellow"].ReplaceRow(redCopy.RotateClockwise(), 0)
	c.Faces["red"].ReplaceCol(whiteCopy.RotateClockwise(), 0)
	c.Faces["green"].RotateClockwise()

	return c
}

func (c *Cube) F2() *Cube {
	return c.F().F()
}

func (c *Cube) Fp() *Cube {
	index := c.CubeSize - 1
	whiteCopy := *(c.Faces["white"].Copy())
	orangeCopy := *(c.Faces["orange"].Copy())
	yellowCopy := *(c.Faces["yellow"].Copy())
	redCopy := *(c.Faces["red"].Copy())
	c.Faces["white"].ReplaceRow(redCopy.RotateAntiClockwise(), index)
	c.Faces["red"].ReplaceCol(yellowCopy.RotateAntiClockwise(), 0)
	c.Faces["yellow"].ReplaceRow(orangeCopy.RotateAntiClockwise(), 0)
	c.Faces["orange"].ReplaceCol(whiteCopy.RotateAntiClockwise(), index)
	c.Faces["green"].RotateAntiClockwise()

	return c
}

func (c *Cube) B() *Cube {
	index := c.CubeSize - 1
	whiteCopy := *(c.Faces["white"].Copy())
	orangeCopy := *(c.Faces["orange"].Copy())
	yellowCopy := *(c.Faces["yellow"].Copy())
	redCopy := *(c.Faces["red"].Copy())
	c.Faces["white"].ReplaceRow(redCopy.RotateAntiClockwise(), 0)
	c.Faces["red"].ReplaceCol(yellowCopy.RotateAntiClockwise(), index)
	c.Faces["yellow"].ReplaceRow(orangeCopy.RotateAntiClockwise(), index)
	c.Faces["orange"].ReplaceCol(whiteCopy.RotateAntiClockwise(), 0)
	c.Faces["blue"].RotateClockwise()

	return c
}

func (c *Cube) B2() *Cube {
	return c.B().B()
}

func (c *Cube) Bp() *Cube {
	index := c.CubeSize - 1
	whiteCopy := *(c.Faces["white"].Copy())
	orangeCopy := *(c.Faces["orange"].Copy())
	yellowCopy := *(c.Faces["yellow"].Copy())
	redCopy := *(c.Faces["red"].Copy())
	c.Faces["white"].ReplaceRow(orangeCopy.RotateClockwise(), 0)
	c.Faces["orange"].ReplaceCol(yellowCopy.RotateClockwise(), 0)
	c.Faces["yellow"].ReplaceRow(redCopy.RotateClockwise(), index)
	c.Faces["red"].ReplaceCol(whiteCopy.RotateClockwise(), index)
	c.Faces["blue"].RotateAntiClockwise()

	return c
}
