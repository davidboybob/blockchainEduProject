/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/davidboybob/blockchainEduProject/cmd"
	"github.com/davidboybob/blockchainEduProject/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
