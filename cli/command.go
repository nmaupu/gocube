package cli

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/nmaupu/gocube/data"
	"os"
)

func Process(appName, appDesc, appVersion string) {
	app := cli.App(appName, appDesc)
	app.Spec = "[-d] -s"

	app.Version("v version", fmt.Sprintf("%s version %s", appName, appVersion))

	var (
		size  = app.IntOpt("s size", 3, "Size of the cube")
		debug = app.BoolOpt("d debug", false, "Enable debug mode")
	)

	app.Action = func() {
		c := data.NewCube(*size)
		data.SetDebug(*debug)
		fmt.Println(c)
	}

	app.Run(os.Args)
}
