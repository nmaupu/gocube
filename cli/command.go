package cli

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/jawher/mow.cli"
	"github.com/nmaupu/gocube/compute"
	"github.com/nmaupu/gocube/cube3d"
	"github.com/nmaupu/gocube/data"
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
	cubieSize = app.IntOpt("c cubie", 30, "Size of one cubie when rendered (px)")
	debug = app.BoolOpt("d debug", false, "Enable debug mode")

	app.Command("scramble", "Scramble with the given algorithm", scramble)
	app.Command("reverse", "Reverse the given algorithm", reverse)
	app.Command("generate", "Generate algs", generate)
	app.Command("exportPDF", "Export as a PDF", exportPDF)
	app.Command("test3D", "Test 3D", test3D)

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

func test3D(cmd *cli.Cmd) {
	cmd.Action = func() {
		c := data.NewCube(*size, float64(*cubieSize))

		ctx := gg.NewContext(1000, 1000)
		ctx.SetHexColor("#FFFFFF")
		ctx.Clear()
		ctx.SetHexColor("#000000")
		ctx.SetLineWidth(1)

		cube3d.DrawCube(ctx, 200, 200, c)

		ctx.SavePNG("/tmp/out.png")
	}
}
