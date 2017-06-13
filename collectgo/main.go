package main

import (
	"os"

	"github.com/drkaka/lg"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "collectgo"
	app.Usage = "Used to collect information."
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Run in debug level.",
		},
		cli.StringFlag{
			Name:  "plugin",
			Value: "plugin",
			Usage: "Path for plugin folder.",
		},
	}

	app.Before = func(c *cli.Context) error {
		lg.InitLogger(c.GlobalBool("debug"))
		lg.L(nil).Debug("logger initialized")
		return nil
	}

	app.Run(os.Args)
}
