package cli

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/jawher/mow.cli"
	"github.com/nmaupu/gocube/compute"
	"github.com/nmaupu/gocube/data"
	"golang.org/x/image/font/gofont/goregular"
	"os"
)

var (
	size      *int
	cubieSize *int
	debug     *bool
)

func Process(appName, appDesc, appVersion string) {
	app := cli.App(appName, appDesc)
	app.Spec = "[-d] [-s] [-c]"

	app.Version("v version", fmt.Sprintf("%s version %s", appName, appVersion))

	size = app.IntOpt("s size", 3, "Size of the cube")
	cubieSize = app.IntOpt("c cubie", 50, "Size of one cubie when rendered (px)")
	debug = app.BoolOpt("d debug", false, "Enable debug mode")

	app.Command("scramble", "Scramble with the given algorithm", scramble)
	app.Command("reverse", "Reverse the given algorithm", reverse)
	app.Command("generate", "Generate algs", generate)
	app.Command("draw", "Drawing tests", draw)
	app.Command("drawOLL", "Draw OLL case", drawOll)

	app.Action = func() {
		c := data.NewCube(*size, float64(*cubieSize))
		data.SetDebug(*debug)
		fmt.Println(c)
	}

	app.Run(os.Args)
}

func scramble(cmd *cli.Cmd) {
	alg := cmd.StringOpt("a alg", "", "Algorithm to use for scrambling")

	cmd.Action = func() {
		c := data.NewCube(*size, float64(*cubieSize))
		data.SetDebug(*debug)

		c.Execute(data.NewAlg(*alg))
		fmt.Println(c)
	}
}

func reverse(cmd *cli.Cmd) {
	cmd.Spec = "-a"
	alg := cmd.StringOpt("a alg", "", "Algorithm to reverse")

	cmd.Action = func() {
		a := data.NewAlg(*alg)
		fmt.Println(a.Reverse())
	}
}

func generate(cmd *cli.Cmd) {
	length := cmd.IntOpt("l length", 20, "Length of the algorithm to generate")
	nb := cmd.IntOpt("n number", 1, "Number of algorithms to generate")
	display := cmd.BoolOpt("d display", false, "Display cube status after scramble")

	cmd.Action = func() {
		g := compute.NewGenerator()
		for i := 0; i < *nb; i++ {
			alg := g.GenerateAlg(*length)
			fmt.Println(alg)
			if *display {
				c := data.NewCube(*size, float64(*cubieSize))
				c.Execute(alg)
				fmt.Print(c)
			}
			fmt.Println("-----")
		}
	}
}

func draw(cmd *cli.Cmd) {
	cmd.Action = func() {
		c := data.NewCube(*size, float64(*cubieSize))
		alg := compute.NewGenerator().GenerateAlg(20)
		fmt.Println(alg)
		c.Execute(alg)
		fmt.Println(c)

		ctx := gg.NewContext(1000, 1000)
		ctx.SetHexColor("#FFFFFF")
		ctx.Clear()
		ctx.SetHexColor("#000000")

		font, err := truetype.Parse(goregular.TTF)
		if err != nil {
			panic("")
		}
		face := truetype.NewFace(font, &truetype.Options{
			Size: 30,
		})
		ctx.SetFontFace(face)

		ctx.DrawString(alg.String(), 10, 40)
		c.Draw(ctx, 0, 100)
		ctx.SavePNG("/tmp/out.png")
	}
}

func drawOll(cmd *cli.Cmd) {
	cmd.Spec = "-a -o [-spc] "
	alg := cmd.StringOpt("a alg", "", "Algorithm corresponding to OLL to display (alg given will be reversed to setup OLL)")
	setup := cmd.StringOpt("s setup", "z2", "Algorithm to use to setup the color face on top, z2 presents yellow on top")
	output := cmd.StringOpt("o output", "", "Output file name (png format)")
	postSetup := cmd.StringOpt("p post", "", "Algorithm to execute at the end to reposition the cube if needed")
	allColors := cmd.BoolOpt("c allColors", false, "Display all colors instead just the one on top")

	cmd.Action = func() {
		// Retrieving the color on top given the setup alg
		c := data.NewCubeColors(*size, float64(*cubieSize), data.Colors)
		setupAlg := data.NewAlg(*setup)
		c.Execute(setupAlg)
		centerColor, err := c.Faces["white"].GetCenterColor()
		if err != nil {
			fmt.Print(err)
		}

		// Generating a cube with the right color for OLL
		a := data.NewAlg(*alg)
		if !*allColors {
			c = data.NewCubeColors(*size, float64(*cubieSize), data.GetColorsOLL(centerColor.Name))
			c.Execute(setupAlg)
		}
		c.Execute(a.Copy().Reverse())
		c.Execute(data.NewAlg(*postSetup))

		fmt.Println(a)
		fmt.Println(c)

		ctx := gg.NewContext(1000, 1000)
		ctx.SetHexColor("#FFFFFF")
		ctx.Clear()
		ctx.SetHexColor("#000000")

		font, err := truetype.Parse(goregular.TTF)
		if err != nil {
			panic("")
		}
		face := truetype.NewFace(font, &truetype.Options{
			Size: 30,
		})
		ctx.SetFontFace(face)

		ctx.DrawString(a.String(), 10, 40)
		// Cube has been z2'ed, so yellow replaced white on data model
		// so displaying white will in fact display yellow
		// for other face, make white replace the color targetted
		c.DrawTopView(ctx, 0, 100, "white")
		ctx.SavePNG(*output)
	}
}
