package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Daily Active CLI"
	app.Usage = "Quickly post what you're doing"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := ""
		if c.NArg() > 0 {
			name = c.Args()[0]
		}
		if c.String("lang") == "italian" {
			fmt.Println("Ciao!", name)
		} else {
			fmt.Println("Hello!")
		}
		return nil
	}
	app.Run(os.Args)
}
