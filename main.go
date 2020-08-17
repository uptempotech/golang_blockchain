package main

import (
	"io/ioutil"
	"os"

	"github.com/uptempotech/golang_blockchain/cli"
	"github.com/uptempotech/golang_blockchain/core"
)

func main() {
	defer os.Exit(0)
	core.InitLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	cmd := cli.CommandLine{}
	cmd.Run()
}
