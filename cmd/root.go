/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "block",
	Short: "This is a tool for BlockChain Education. 🌎💲🍖🥙",
	Long: `This is a tool for BlockChain Education. 🌎💲🍖🥙

 Hi, We're Developer Seong Jin, Kyun Tae to be BlockChain Experts.
 💸 This app is education for BlockChain Tech.
 💰 This app will help you with your development skills.
 💶 This app is divided into levels so that you can use it accoding to your abilities.
	
 If you like our CLI APP, you can buy a cup of coffee to us. ☕

 Contact to us 📩 davidboybob7780@gmail.com, 📩 kore17641764@gmail.com. 
	
	
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.blockchainEduProject.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
