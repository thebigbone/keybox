package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

const (
	CheckSymbol = "\u2714 "
	CrossSymbol = "\u2716 "
)

func initSSHKeysStore(con *cli.Context) error {
	keyStorePath, err := getPath()

	_, err = os.Stat(keyStorePath)

	if err == nil {
		color.Red("%s keystore folder already exists!", CrossSymbol)
		return nil
	} else if !os.IsNotExist(err) {
		log.Fatal(err)
	}

	err = os.MkdirAll(keyStorePath, os.ModePerm)

	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(color.GreenString(fmt.Sprintf("%s keystore created at %s", CheckSymbol, keyStorePath)))
	return nil
}
