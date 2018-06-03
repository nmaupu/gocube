package data

import (
	"bytes"
	"strings"
)

type Cube struct {
	Faces    map[string]Face
	CubeSize int
	moves    map[string](func() *Cube)
}

func NewCube(cubeSize int) *Cube {
	c := Cube{
		Faces:    make(map[string]Face),
		CubeSize: cubeSize,
	}

	// Creating moves funcs map
	// So that we can make a call like:
	// c.moves["R'"]()
	c.moves = make(map[string](func() *Cube))
	// Common moves
	c.moves["R"] = c.R
	c.moves["L"] = c.L
	c.moves["F"] = c.F
	c.moves["B"] = c.B
	c.moves["D"] = c.D
	c.moves["U"] = c.U
	c.moves["M"] = c.M
	c.moves["S"] = c.S
	c.moves["E"] = c.E

	// Reverse moves
	c.moves["R'"] = c.Rp
	c.moves["L'"] = c.Lp
	c.moves["F'"] = c.Fp
	c.moves["B'"] = c.Bp
	c.moves["D'"] = c.Dp
	c.moves["U'"] = c.Up
	c.moves["M'"] = c.Mp
	c.moves["S'"] = c.Sp
	c.moves["E'"] = c.Ep

	// Double moves
	c.moves["R2"] = c.R2
	c.moves["L2"] = c.L2
	c.moves["F2"] = c.F2
	c.moves["B2"] = c.B2
	c.moves["D2"] = c.D2
	c.moves["U2"] = c.U2
	c.moves["M2"] = c.M2
	c.moves["S2"] = c.S2
	c.moves["E2"] = c.E2

	// Rotation moves
	c.moves["x"] = c.X
	c.moves["y"] = c.Y
	c.moves["z"] = c.Z
	c.moves["x'"] = c.Xp
	c.moves["y'"] = c.Yp
	c.moves["z'"] = c.Zp
	c.moves["x2"] = c.X2
	c.moves["y2"] = c.Y2
	c.moves["z2"] = c.Z2

	// Wide moves
	c.moves["r"] = c.Rw
	c.moves["r'"] = c.Rwp
	c.moves["r2"] = c.Rw2
	c.moves["l"] = c.Lw
	c.moves["l'"] = c.Lwp
	c.moves["l2"] = c.Lw2
	c.moves["d"] = c.Dw
	c.moves["d'"] = c.Dwp
	c.moves["d2"] = c.Dw2
	c.moves["u"] = c.Uw
	c.moves["u'"] = c.Uwp
	c.moves["u2"] = c.Uw2
	c.moves["b"] = c.Bw
	c.moves["b'"] = c.Bwp
	c.moves["b2"] = c.Bw2
	c.moves["f"] = c.Fw
	c.moves["f'"] = c.Fwp
	c.moves["f2"] = c.Fw2
	c.moves["m"] = c.Mw
	c.moves["m'"] = c.Mwp
	c.moves["m2"] = c.Mw2
	c.moves["s"] = c.Sw
	c.moves["s'"] = c.Swp
	c.moves["s2"] = c.Sw2
	c.moves["e"] = c.Ew
	c.moves["e'"] = c.Ewp
	c.moves["e2"] = c.Ew2

	// Creating faces
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

func (c *Cube) M() *Cube {
	mid := int(c.CubeSize / 2)
	blueCopy := *(c.Faces["blue"].Copy())
	yellowCopy := *(c.Faces["yellow"].Copy())
	whiteCopy := *(c.Faces["white"].Copy())
	c.Faces["white"].ReplaceCol(blueCopy.RotateClockwise().RotateClockwise(), mid)
	c.Faces["blue"].ReplaceCol(yellowCopy.RotateClockwise().RotateClockwise(), mid)
	c.Faces["yellow"].ReplaceCol(c.Faces["green"], mid)
	c.Faces["green"].ReplaceCol(whiteCopy, mid)

	return c
}

func (c *Cube) Mp() *Cube {
	return c.M().M().M()
}

func (c *Cube) M2() *Cube {
	return c.M().M()
}

func (c *Cube) S() *Cube {
	mid := int(c.CubeSize / 2)
	whiteCopy := *(c.Faces["white"].Copy())
	orangeCopy := *(c.Faces["orange"].Copy())
	yellowCopy := *(c.Faces["yellow"].Copy())
	redCopy := *(c.Faces["red"].Copy())
	c.Faces["white"].ReplaceRow(orangeCopy.RotateClockwise(), mid)
	c.Faces["orange"].ReplaceCol(yellowCopy.RotateClockwise(), mid)
	c.Faces["yellow"].ReplaceRow(redCopy.RotateClockwise(), mid)
	c.Faces["red"].ReplaceCol(whiteCopy.RotateClockwise(), mid)

	return c
}

func (c *Cube) Sp() *Cube {
	return c.S().S().S()
}

func (c *Cube) S2() *Cube {
	return c.S().S()
}

func (c *Cube) E() *Cube {
	return c.Z().M().Zp()
}

func (c *Cube) Ep() *Cube {
	return c.E().E().E()
}

func (c *Cube) E2() *Cube {
	return c.E().E()
}

func (c *Cube) X() *Cube {
	whiteCopy := *(c.Faces["white"].Copy())

	c.Faces["white"] = c.Faces["green"]
	c.Faces["green"] = c.Faces["yellow"]
	c.Faces["orange"].RotateAntiClockwise()
	c.Faces["red"].RotateClockwise()
	c.Faces["yellow"] = c.Faces["blue"].Flip()
	c.Faces["blue"] = whiteCopy.Flip()

	return c
}

func (c *Cube) Y() *Cube {
	copyGreen := *(c.Faces["green"].Copy())

	c.Faces["white"].RotateClockwise()
	c.Faces["yellow"].RotateAntiClockwise()
	c.Faces["green"] = c.Faces["red"]
	c.Faces["red"] = c.Faces["blue"]
	c.Faces["blue"] = c.Faces["orange"]
	c.Faces["orange"] = copyGreen

	return c
}

func (c *Cube) Z() *Cube {
	whiteCopy := *(c.Faces["white"].Copy())
	c.Faces["white"] = c.Faces["orange"].RotateClockwise()
	c.Faces["orange"] = c.Faces["yellow"].RotateClockwise()
	c.Faces["yellow"] = c.Faces["red"].RotateClockwise()
	c.Faces["red"] = whiteCopy.RotateClockwise()
	c.Faces["green"].RotateClockwise()
	c.Faces["blue"].RotateAntiClockwise()

	return c
}

func (c *Cube) Xp() *Cube {
	return c.X().X().X()
}

func (c *Cube) Yp() *Cube {
	return c.Y().Y().Y()
}

func (c *Cube) Zp() *Cube {
	return c.Z().Z().Z()
}

func (c *Cube) X2() *Cube {
	return c.X().X()
}

func (c *Cube) Y2() *Cube {
	return c.Y().Y()
}

func (c *Cube) Z2() *Cube {
	return c.Z().Z()
}

func (c *Cube) Rw() *Cube {
	return c.R().Mp()
}

func (c *Cube) Rwp() *Cube {
	return c.Rp().M()
}

func (c *Cube) Rw2() *Cube {
	return c.Rw().Rw()
}

func (c *Cube) Lw() *Cube {
	return c.L().M()
}

func (c *Cube) Lwp() *Cube {
	return c.Lp().Mp()
}

func (c *Cube) Lw2() *Cube {
	return c.Lw().Lw()
}

func (c *Cube) Uw() *Cube {
	return c.U().Ep()
}

func (c *Cube) Uwp() *Cube {
	return c.Up().E()
}

func (c *Cube) Uw2() *Cube {
	return c.Uw().Uw()
}

func (c *Cube) Dw() *Cube {
	return c.D().E()
}

func (c *Cube) Dwp() *Cube {
	return c.Dp().Ep()
}

func (c *Cube) Dw2() *Cube {
	return c.Dw().Dw()
}

func (c *Cube) Fw() *Cube {
	return c.F().S()
}

func (c *Cube) Fwp() *Cube {
	return c.Fp().Sp()
}

func (c *Cube) Fw2() *Cube {
	return c.Fw().Fw()
}

func (c *Cube) Bw() *Cube {
	return c.B().Sp()
}

func (c *Cube) Bwp() *Cube {
	return c.Bp().S()
}

func (c *Cube) Bw2() *Cube {
	return c.Bw().Bw()
}

func (c *Cube) Mw() *Cube {
	return c.L().Rp()
}

func (c *Cube) Mwp() *Cube {
	return c.Lp().R()
}

func (c *Cube) Mw2() *Cube {
	return c.Mw().Mw()
}

func (c *Cube) Sw() *Cube {
	return c.F().Bp()
}

func (c *Cube) Swp() *Cube {
	return c.Fp().B()
}

func (c *Cube) Sw2() *Cube {
	return c.Sw().Sw()
}

func (c *Cube) Ew() *Cube {
	return c.Up().D()
}

func (c *Cube) Ewp() *Cube {
	return c.U().Dp()
}

func (c *Cube) Ew2() *Cube {
	return c.Ew().Ew()
}

func (c *Cube) Execute(alg Alg) *Cube {
	for _, m := range alg.Moves {
		// Getting and calling corresponding func
		// From funcs map
		c.moves[m]()
	}

	return c
}
