package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"bufio"
)

var rootCmd = &cobra.Command{
	Use: "ccwc",
	Short: "ccwc is a CLI tool for the Code Challenge Workshop",
	Long: `ccwc is a CLI tool for the Code Challenge Workshop. It is a collection of challenges to help you improve your programming skills.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(os.Args) < 3 {
			fmt.Println("Please provide flag for how to count the file.")
			return
		}

		if len(os.Args) < 4 {
			fmt.Println("Please provide a file name as a command-line argument.")
			return
		}

		filename := os.Args[3]
		file, err := os.Open(filename)

		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		defer file.Close()

		if (cmd.Flags().Lookup("c").Changed) {
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanBytes)

			count := 0
			for scanner.Scan() {
				count++
			}

			fmt.Println(count, filename)
		} else if (cmd.Flags().Lookup("w").Changed) {
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanWords)

			count := 0
			for scanner.Scan() {
				count++
			}

			fmt.Println(count, filename)
		} else if (cmd.Flags().Lookup("l").Changed) {
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)

			count := 0
			for scanner.Scan() {
				count++
			}

			fmt.Println(count, filename)
		} else if (cmd.Flags().Lookup("m").Changed) {
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanRunes)

			count := 0
			for scanner.Scan() {
				count++
			}

			fmt.Println(count, filename)
		}
	},
}

var (
	flag string
)

func Execute() {
	rootCmd.Flags().StringVarP(&flag, "c", "c", "", "count the number of bytes in a text file")
	rootCmd.Flags().StringVarP(&flag, "w", "w", "", "count the number of words in a text file")
	rootCmd.Flags().StringVarP(&flag, "l", "l", "", "count the number of lines in a text file")
	rootCmd.Flags().StringVarP(&flag, "m", "m", "", "count the number of characters in a text file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}