/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "block",
	Short: "Block chain Education app!",
	Long:  `How to learn Blockchain? You can learn Blockchian with us! `,
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			return printVersion()
		}
		return cmd.Help()
	},
}

func printVersion() error {
	const version = "1.0"
	fmt.Println("Version: ", version)
	return nil
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
