package main

import (
	"github.com/urfave/cli"
	"go-cli/cmd"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "apcli"
	app.Usage = "Auto make Apache+PHP"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:   "all_update",
			Action: cmd.All_Update,
			Usage:  "system update",
		},
		{
			Name:   "cli-apache",
			Action: cmd.Setup_apa,
			Usage:  "setup apache",
		},
		{
			Name:   "cli-php",
			Action: cmd.Setup_php,
			Usage:  "setup php",
		},
	}
	app.Run(os.Args)
}
