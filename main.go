package main

import (
	"os"
	"github.com/codegangsta/cli"
	"yoseage/yose"
)

func main() {
	app := cli.NewApp()
	app.Name = "yoseage"
	app.Usage = "compressor file"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) {
		var dirPaths = []string{"/root/document"}
		yose.CheckTarget(dirPaths)
	}
	app.Run(os.Args)
}
