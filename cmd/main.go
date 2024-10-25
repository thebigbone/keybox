package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "keybox",
		Usage: "cli for managing your ssh keys",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "init SSH keys store for the first time use",
				Action: func(con *cli.Context) error {
					return initSSHKeysStore(con)
				},
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "create SSH key in keystore",
				Action: func(con *cli.Context) error {
					return createSSHKeys(con)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
