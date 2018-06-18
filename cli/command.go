package cli

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/jawher/mow.cli"
	"github.com/nmaupu/gocube/compute"
	"github.com/nmaupu/gocube/cube3d"
	"github.com/nmaupu/gocube/data"
	"image"
	"image/png"
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
	output := cmd.StringOpt("o output", "/tmp/out.png", "Output file name")

	cmd.Action = func() {
		c := data.NewCube(*size, float64(*cubieSize))
		//g := compute.NewGenerator()
		//alg := g.GenerateAlg(20)
		//fmt.Println(alg)
		c.Execute(data.NewAlg("U' B' D2 L2 F2 R2 U B' L2 B' U B' D F2 U2 D2 B D U R2"))

		imgDim := 700
		ctx := gg.NewContext(imgDim, imgDim)
		ctx.SetHexColor("#FFFFFF")
		ctx.Clear()
		ctx.SetHexColor("#000000")
		ctx.SetLineWidth(1)

		cube3d.DrawCube(ctx, c)

		ctx.SavePNG(*output)

		cropImage(*output)
	}
}

func cropImage(in string) error {
	f, _ := os.Open(in)
	img, _, _ := image.Decode(f)

	// Store all non alpha coords
	xs := make([]int, 0)
	ys := make([]int, 0)

	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			px := img.At(x, y)
			r, g, b, _ := px.RGBA()
			if r != 65535 && g != 65535 && b != 65535 {
				xs = append(xs, x)
				ys = append(ys, y)
			}
		}
	}

	xL, _ := getMin(xs, ys)
	xR, _ := getMax(xs, ys)
	yU, _ := getMin(ys, xs)
	yD, _ := getMax(ys, xs)

	fso, err := os.Create("/tmp/test.png")
	if err != nil {
		return err
	}
	defer fso.Close()

	croppedimg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(xL, yU, xR, yD))

	return png.Encode(fso, croppedimg)
}

func getMin(xs, ys []int) (int, int) {
	min := 100000000
	y := 0
	for k, v := range xs {
		if v < min {
			min = v
			y = ys[k]
		}
	}

	return min, y
}
func getMax(xs, ys []int) (int, int) {
	max := 0
	y := 0
	for k, v := range xs {
		if v > max {
			max = v
			y = ys[k]
		}
	}

	return max, y
}
