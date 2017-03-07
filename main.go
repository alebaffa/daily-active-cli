package main

import (
	"log"
	"os"

	"github.com/alebaffa/daily-cli/cmd"
	"github.com/urfave/cli"
)

const (
	descAction   = "Write what you are doing."
	descFlagType = "the type of activity (coding, reading, studying, etc)."
	descFlagDesc = "description of the activity."
	descFlagURL  = "URL related to the activity."
)

func main() {
	app := cli.NewApp()
	app.Name = "daily-active-cli"
	app.Usage = "Quickly update your Daily Active Log from the terminal."
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		{
			Name:  "action",
			Usage: descAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "type, t",
					Usage: descFlagType,
				},
				cli.StringFlag{
					Name:  "desc, d",
					Usage: descFlagDesc,
				},
				cli.StringFlag{
					Name:  "url",
					Usage: descFlagURL,
				},
			},
			Action: cmd.Action,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
