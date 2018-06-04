package cli

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/nmaupu/gocube/compute"
	"github.com/nmaupu/gocube/data"
	"os"
)

var (
	size  *int
	debug *bool
)

func Process(appName, appDesc, appVersion string) {
	app := cli.App(appName, appDesc)
	app.Spec = "[-d] -s"

	app.Version("v version", fmt.Sprintf("%s version %s", appName, appVersion))

	size = app.IntOpt("s size", 3, "Size of the cube")
	debug = app.BoolOpt("d debug", false, "Enable debug mode")

	app.Command("scramble", "Scramble with the given algorithm", scramble)
	app.Command("reverse", "Reverse the given algorithm", reverse)
	app.Command("generate", "Generate algs", generate)

	app.Action = func() {
		c := data.NewCube(*size)
		data.SetDebug(*debug)
		fmt.Println(c)
	}

	app.Run(os.Args)
}

func scramble(cmd *cli.Cmd) {
	alg := cmd.StringOpt("a alg", "", "Algorithm to use for scrambling")

	cmd.Action = func() {
		c := data.NewCube(*size)
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
				c := data.NewCube(*size)
				c.Execute(alg)
				fmt.Print(c)
			}
			fmt.Println("-----")
		}
	}
}
