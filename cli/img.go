package cli

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/jawher/mow.cli"
	"github.com/nmaupu/gocube/data"
	"github.com/nmaupu/gocube/image"
)

func exportImg(cmd *cli.Cmd) {
	cmd.Spec = "-o [-t] [-i...] [-a] [-f] [-p]"
	outputArg := cmd.StringOpt("o output", "", "Image file name (png)")
	viewArg := cmd.StringOpt("t viewType", "full", "drawing style (3d, top, full)")
	colorsArg := cmd.StringsOpt("i include", nil, "Include color (colors: yellow, white, green, blue, red, orange) - default: all colors are included")
	algArg := cmd.StringOpt("a alg", "", "Algorithm to execute")
	preAlgArg := cmd.StringOpt("p preAlg", "", "Algorithm to execute to setup up the cube (use x, y and z to position the needed cube colors on top and sides)")
	f2lMode := cmd.BoolOpt("f f2l", false, "F2L mode (display only centers and displayed corner pieces) - if this option is activated, viewType is set to 3d")

	cmd.Action = func() {
		if (*viewArg != "full" && *viewArg != "3d" && *viewArg != "top") ||
			*f2lMode {
			fmt.Println("Auto switching to 3d view mode")
			*viewArg = "3d"
		}
		if *colorsArg == nil {
			fmt.Println("All colors are included")
			colorsArg = &[]string{
				"white", "yellow",
				"green", "blue",
				"red", "orange",
			}
		}

		alg := data.NewAlg(*algArg)
		preAlg := data.NewAlg(*preAlgArg)
		var colors map[string]data.Color
		colors = data.GetColors(*colorsArg...)

		c := data.NewCubeColors(*size, float64(cubieSize), colors)
		c.Execute(preAlg)
		if *f2lMode {
			fmt.Println("F2L mode activated")
			c.F2lColorMode()
		}
		c.Execute(alg)

		var ctx *gg.Context
		switch *viewArg {
		case "top":
			ctx = c.DrawTopView("white")
		case "full":
			ctx = c.Draw()
		case "3d":
			ctx = c.Draw3d(1000)
			ctx, _ = image.TrimImageWhite(ctx)
		default:
			panic(fmt.Sprintf("Incorrect view, %s", viewArg))
		}

		fmt.Println("Writing png file to:", *outputArg)
		ctx.SavePNG(*outputArg)
	}
}
