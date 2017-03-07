package cmd

import (
	"log"
	"os"

	"github.com/alebaffa/daily-cli/xml"
	"github.com/urfave/cli"
)

func Action(ctx *cli.Context) error {
	t := ctx.String("type")
	d := ctx.String("desc")
	url := ctx.String("url")

	if t == "" || d == "" {
		log.Fatal("You need both options -type and -desc")
	}

	output := xml.EncodeXML(t, d, url)
	os.Stdout.Write(output)
	return nil
}
