package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func initSSHKeysStore(con *cli.Context) error {
	keyStorePath, err := getPath()
	_, err = os.Stat(keyStorePath)
	if os.IsNotExist(err) {
		log.Fatal("keystore folder already exists!")
	}

	err = os.MkdirAll(keyStorePath, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
