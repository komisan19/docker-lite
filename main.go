package main

import (
	"github.com/urfave/cli"
	"go-cli/cmd"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "docker-like"
	app.Usage = "lite docker-cli"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:   "version",
			Action: cmd.Version,
			Usage:  "docker version",
		},
		{
			Name:   "ps",
			Action: cmd.Ps,
			Usage:  "docker ps",
		},
	}
	app.Run(os.Args)
}
