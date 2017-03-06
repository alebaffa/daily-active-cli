package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

func Action(ctx *cli.Context) error {
	text := ctx.String("text")
	url := ctx.String("url")

	fmt.Println(text + " and " + url)
	return nil
}
