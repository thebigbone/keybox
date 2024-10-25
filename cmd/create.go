package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func createSSHKeys(con *cli.Context) error {
	keyName := con.Args().First()

	fmt.Println(keyName)
	return nil
}
