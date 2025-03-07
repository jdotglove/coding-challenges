package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "ccwc",
	Short: "A CLI tool for the count number of words",
	Long: "CCWC is a CLI tool that counts the number of words in a given text file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing CCWC '%s'\n", err)
		os.Exit(1)
	}
}