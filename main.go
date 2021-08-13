package main

import (
	"github.com/presuit/nomadcoin/cli"
	"github.com/presuit/nomadcoin/db"
)

func main() {
	cli.Start()
	defer db.Close()
}
