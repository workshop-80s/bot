package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"bot/cmd/article"
	"bot/cmd/subscribe"
)

var rootCmd = &cobra.Command{
	Use:   "bot",
	Short: "",
	Long: `This is my bot`,
	Run: func(cmd *cobra.Command, args []string) { 
		fmt.Println("welcome to my bot")
	},
}

// Execute executes the root command.
func Execute() error {
	article.AddCommand(rootCmd)
	subscribe.AddCommand(rootCmd)
	return rootCmd.Execute()
}

