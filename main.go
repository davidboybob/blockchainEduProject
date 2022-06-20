/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"block/cmd"
	"block/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
