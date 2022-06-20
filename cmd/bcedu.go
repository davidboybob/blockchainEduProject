/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// blockchainCmd represents the blockchain command
var bceduCmd = &cobra.Command{
	Use:   "bcedu",
	Short: "blockchian education cli",
	Long:  `Block chain Education!!`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()

	},
}

func init() {
	rootCmd.AddCommand(bceduCmd)

}
