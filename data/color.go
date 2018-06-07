package data

import (
	"github.com/fogleman/gg"
)

type Color struct {
	Color     int
	Name      string
	ShortName string
	Debug     string
	HexColor  string
}

func (c Color) String() string {
	if !GetDebug() {
		return c.ShortName
	} else {
		return c.Debug
	}
}

func (c Color) Draw(ctx *gg.Context, x float64, y float64, cubieSize float64) {
	border := cubieSize * 6 / 100
	rounded := cubieSize * 5 / 100
	margin := cubieSize * 1 / 100

	ctx.Push()

	ctx.DrawRectangle(x, y, cubieSize, cubieSize)
	ctx.SetRGB255(255, 255, 255)
	ctx.Fill()

	ctx.DrawRoundedRectangle(x+margin, y+margin, cubieSize-margin*2, cubieSize-margin*2, rounded)
	ctx.SetRGB255(0, 0, 0)
	ctx.Fill()

	ctx.DrawRoundedRectangle(
		x+margin+border,
		y+margin+border,
		cubieSize-border*2-margin*2,
		cubieSize-border*2-margin*2,
		rounded)
	ctx.SetHexColor(c.HexColor)
	ctx.Fill()

	ctx.Pop()
}

// Draw width x half height cubie - Horizontally
func (c Color) DrawHalfH(ctx *gg.Context, x float64, y float64, cubieSize float64) {
	border := cubieSize * 6 / 100
	rounded := cubieSize * 5 / 100
	margin := cubieSize * 1 / 100

	ctx.Push()

	ctx.DrawRectangle(x, y, cubieSize, cubieSize/2)
	ctx.SetRGB255(255, 255, 255)
	ctx.Fill()

	ctx.DrawRoundedRectangle(x+margin, y+margin, cubieSize-margin*2, (cubieSize/2)-margin*2, rounded)
	ctx.SetRGB255(0, 0, 0)
	ctx.Fill()

	ctx.DrawRoundedRectangle(
		x+margin+border,
		y+margin+border,
		cubieSize-border*2-margin*2,
		(cubieSize/2)-border*2-margin*2,
		rounded)
	ctx.SetHexColor(c.HexColor)
	ctx.Fill()

	ctx.Pop()
}

// Draw width x half height cubie - Vertically
func (c Color) DrawHalfV(ctx *gg.Context, x float64, y float64, cubieSize float64) {
	border := cubieSize * 6 / 100
	rounded := cubieSize * 5 / 100
	margin := cubieSize * 1 / 100

	ctx.Push()

	ctx.DrawRectangle(x, y, cubieSize/2, cubieSize)
	ctx.SetRGB255(255, 255, 255)
	ctx.Fill()

	ctx.DrawRoundedRectangle(x+margin, y+margin, (cubieSize/2)-margin*2, cubieSize-margin*2, rounded)
	ctx.SetRGB255(0, 0, 0)
	ctx.Fill()

	ctx.DrawRoundedRectangle(
		x+margin+border,
		y+margin+border,
		(cubieSize/2)-border*2-margin*2,
		cubieSize-border*2-margin*2,
		rounded)
	ctx.SetHexColor(c.HexColor)
	ctx.Fill()

	ctx.Pop()
}
