package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"bufio"
	"regexp"
)

var rootCmd = &cobra.Command{
	Use: "jparse",
	Short: "jparse is a CLI tool for the Code Challenge Workshop which can be used to parse JSON files",
	Long: `
		jparse is a CLI tool for the Code Challenge Workshop. 
		It is a collection of challenges to help you improve your programming skills. 
		This tool can be used to parse JSON files.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		filename := os.Args[2]
		validJson := ValidateJson(filename)
		if (validJson == 1) {
			fmt.Println("Invalid JSON")
		} else {
			fmt.Println("Valid JSON")
		}
	},
}

var (
	flag string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ValidateJson(args ...string) int {
	file, err := os.Open(args[0])

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return 0
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	var objectStart bool = false
	var objectEnd bool = false
	var valueStart bool = false
	var valueEnd bool = false
	var foundKey bool = false
	var keyStart bool = false
	var trailingComma bool = false

	for scanner.Scan() {
		if (scanner.Text() == "{" && !objectStart) {
			objectStart = true
		}

		if (scanner.Text() == "}" && objectStart && !objectEnd) {
			objectEnd = true
		}

		if (scanner.Text() != "{" && scanner.Text() != "}" && objectStart) {
			match, _ := regexp.MatchString("[a-zA-Z]", scanner.Text())
			if (scanner.Text() == "\"") {
				if (!valueStart) {
					trailingComma = false
					valueStart = true
					if (!foundKey && !keyStart) {
						keyStart = true
					}
				} else if (valueStart && !valueEnd) {
					valueStart = false

					if (keyStart) {
						keyStart = false
						foundKey = true
					}
				}
			} else if (match) {
				if (match && !valueStart) {
					return 1
				}
			} else if (!valueStart && scanner.Text() == ":") {
				if (!foundKey) {
					return 1
				}
			} else if (scanner.Text() == "\n") {
				foundKey = false
				valueStart = false
				valueEnd = false
				keyStart = false
			} else if (scanner.Text() == ",") {
				trailingComma = true
			}
		}
	}

	if ((!objectStart && !objectEnd) || trailingComma) {
		return 1
	}

	return 0
}