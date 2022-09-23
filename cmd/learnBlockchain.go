package cmd

import (
	"fmt"
	"os"

	"github.com/davidboybob/blockchainEduProject/data"
	"github.com/davidboybob/blockchainEduProject/global"
	"github.com/manifoldco/promptui"
)

func learnBlockchainSelect() string {
	prompt := promptui.Select{
		Label: "Select Mode",
		Items: []string{"📄Level Test",
			"😁Learning for me"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	switch result {
	case "📄Level Test":
		CreateQuestionBank(global.LevelTestQuestionCount)
		learnBlockchainSelect()
	case "😁Learning for me":
		score := data.GetUserScore()
		level := levelConverter(int(score))
		fmt.Println("Your score is ", score, "Let's start ", level, " Lecture!")
		switch level {
		case global.Beginner:
			fmt.Println("go beginner step")
		case global.Intermediate:
			fmt.Println("go intermediate step")
		case global.Advanced:
			fmt.Println("go advanced step")
		}
	}

	return result
}
