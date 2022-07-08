/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Eduation Start",
	Long:  `Start Education`,
	Run: func(cmd *cobra.Command, args []string) {
		promptStartSelect()
	},
}

// type promptContents struct {
// 	errorMsg string
// 	label    string
// }

func promptStartSelect() string {
	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"chap1_Concept",
			"chap2_Create Blockchain",
			"chap3_Create DB",
			"chap4_Create Wallet",
			"chap5_P2P"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	switch result {
	case "chap1_Concept":
		fmt.Println("Choose chap1_Concept")
		// 함수
	case "chap2_Create Blockchain":
		fmt.Println("Choose Chap2_Create Blockchain")
	case "chap3_Create DB":
		fmt.Println("Choose Chap3_Create DB")
	case "chap4_Create Wallet":
		fmt.Println("Choose Chap4_Create Wallet")
	case "chap5_P2P":
		fmt.Println("Choose chap5_P2P")
	}

	return result
}

func init() {
	rootCmd.AddCommand(startCmd)
}
