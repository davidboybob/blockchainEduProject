package learn

import (
	"fmt"
	"os"

	// "github.com/davidboybob/blockchainEduProject/cmd"
	"github.com/davidboybob/blockchainEduProject/global"
	"github.com/manifoldco/promptui"
)

func LearnForBeginnerSelect() string {
	templates := &promptui.SelectTemplates{
		Label: "{{ . | green }}",
		Help:  "It is a category that summarizes blockchain terms for beginners.",
	}

	prompt := promptui.Select{
		Label:     "Select Mode",
		Items:     []string{"📜 Terms"},
		Templates: templates,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	switch result {
	case "📜 Terms":
		fmt.Println("Terms of Blockchain")
		// cmd.LearnBlockchainSelect()
		for key, val := range global.BeginnerTerms {
			fmt.Printf("%s: %s\n", key, val)
		}
	}

	return result
}
