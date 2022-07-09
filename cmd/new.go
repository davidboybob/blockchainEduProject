/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/davidboybob/blockchainEduProject/data"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Cretae a new blockchian edu note",
	Long:  `Cretae a new blockchian edu note`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()
	},
}

// type promptContent struct {
// 	errorMsg string
// 	label    string
// }

func init() {
	// bceduCmd.AddCommand(newCmd)
}

// func promptGetInput(pc promptContent) string {
// 	validate := func(input string) error {
// 		if len(input) <= 0 {
// 			return errors.New(pc.errorMsg)
// 		}
// 		return nil
// 	}

// 	templates := &promptui.PromptTemplates{
// 		Prompt:  "{{ . }}",
// 		Valid:   "{{ . | green}}",
// 		Invalid: "{{ . | red}}",
// 		Success: "{{ . | bold}}",
// 	}

// 	prompt := promptui.Prompt{
// 		Label:     pc.label,
// 		Templates: templates,
// 		Validate:  validate,
// 	}

// 	result, err := prompt.Run()
// 	if err != nil {
// 		fmt.Printf("Prompt failed %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("Input: %s\n", result)
// 	return result
// }

func promptGetSelect(pc promptContent) string {
	items := []string{"step1", "step2", "step3", "step4"}
	index := -1

	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}
		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Input: %s \n", result)
	return result
}

func createNewNote() {
	wordPromptContent := promptContent{
		"Please provide a step",
		"What step do you want?",
	}

	word := promptGetInput(wordPromptContent)

	definitionPromptContent := promptContent{
		"Please provide a definition",
		fmt.Sprintf("Waht is the definition of %s?\n", word),
	}
	definition := promptGetInput(definitionPromptContent)

	categoryPromptContent := promptContent{
		"Please provide a category",
		fmt.Sprintf("What category does %s belong to?", word),
	}
	category := promptGetSelect(categoryPromptContent)

	data.InsertNote(word, definition, category)
}
