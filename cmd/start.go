/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/davidboybob/blockchainEduProject/data"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Eduation Start",
	Long:  `Start Education`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("start called")

		promptStartSelect()
	},
}

type userInfo struct {
	name           string
	tutorial_score int
	last_step      string
}

func createNewUser() {
	// User 에 데이터 있는 지 확인
	log.Println("displayalluser")
	data.DisplayAllUser()

}

// type promptContents struct {
// 	errorMsg string
// 	label    string
// }
func promptChapSelect() string {
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

func promptStartSelect() string {
	prompt := promptui.Select{
		Label: "Select Mode",
		Items: []string{"Nomad Coin Lecture",
			"Block Chain level Test",
			"Learn Block Chain"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	switch result {
	case "Nomad Coin Lecture":
		fmt.Println("Nomad Coin Lecture")
	case "Block Chain level Test":
		fmt.Println("Block Chain level Test")
		createNewUser()

	case "Learn Block Chain":
		fmt.Println("Learn Block Chain")
		promptChapSelect()
	}

	return result
}

func init() {
	rootCmd.AddCommand(startCmd)
}
