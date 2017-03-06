package main

import (
	"os"

	"github.com/alebaffa/daily-cli/cmd"
	"github.com/urfave/cli"
)

const (
	descAction = "Write what you are doing."
)

func main() {
	app := cli.NewApp()
	app.Name = "daily-active-cli"
	app.Usage = "Quickly update your Daily Active Log from the terminal"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		{
			Name:  "action",
			Usage: descAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "text",
					Usage: "description of your current activity",
				},
				cli.StringFlag{
					Name:  "url",
					Usage: "URL related to your current activity",
				},
			},
			Action: cmd.Action,
		},
	}

	app.Run(os.Args)
}
