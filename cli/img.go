package cli

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/jawher/mow.cli"
	"github.com/nmaupu/gocube/data"
	"github.com/nmaupu/gocube/image"
	"log"
)

func exportImg(cmd *cli.Cmd) {
	cmd.Spec = "-o [-w] [-t] [-i...] [-p] [-a]"
	outputArg := cmd.StringOpt("o output", "", "Image file name (png)")
	widthArg := cmd.IntOpt("w width", 500, "Generated image width")
	viewArg := cmd.StringOpt("t viewType", "full", "drawing style (3d, top, full, f2l)")
	colorsArg := cmd.StringsOpt("i include", nil, "Include color (colors: yellow, white, green, blue, red, orange) - default: all colors are included")
	algArg := cmd.StringOpt("a alg", "", "Algorithm to execute")
	preAlgArg := cmd.StringOpt("p preAlg", "", "Algorithm to execute to setup up the cube (use x, y and z to position the needed cube colors on top and sides)")

	cmd.Action = func() {
		if *viewArg != "full" && *viewArg != "3d" && *viewArg != "top" && *viewArg != "f2l" {
			log.Println("Auto switching to 3d view mode")
			*viewArg = "3d"
		}
		if *colorsArg == nil {
			log.Println("All colors are included")
			colorsArg = &[]string{
				"white", "yellow",
				"green", "blue",
				"red", "orange",
			}
		}
		// Max size is 2000
		if *widthArg > 2000 {
			*widthArg = 2000
		}

		alg := data.NewAlg(*algArg)
		preAlg := data.NewAlg(*preAlgArg)
		var colors map[string]data.Color
		colors = data.GetColors(*colorsArg...)

		c := data.NewCubeColors(*size, float64(cubieSize), colors)
		c.Execute(preAlg)
		if *viewArg == "f2l" {
			err := c.F2lColorMode()
			if err != nil {
				log.Fatalf("Error: %s", err)
			}
		}
		c.Execute(alg)

		var ctx *gg.Context
		switch *viewArg {
		case "top":
			ctx = c.DrawTopView("white")
		case "full":
			ctx = c.Draw()
		case "3d", "f2l":
			size := 1000
			if *widthArg > size {
				size = *widthArg
			}
			ctx = c.Draw3d(size)
			ctx, _ = image.TrimImageWhite(ctx)
		default:
			// Should never happen
			panic(fmt.Sprintf("Incorrect view, %s", *viewArg))
		}

		dstImg := imaging.Resize(ctx.Image(), *widthArg, 0, imaging.Lanczos)
		ctx = gg.NewContextForImage(dstImg)
		log.Println("Writing png file to:", *outputArg)
		ctx.SavePNG(*outputArg)

	}
}
