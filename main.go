package main

import (
	"github.com/AradTenenbaum/GoBlockchain/blockchain"
	"github.com/AradTenenbaum/GoBlockchain/cli"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.Db.Close()

	cli := cli.CLI{
		Bc: bc,
	}
	cli.Run()
}
