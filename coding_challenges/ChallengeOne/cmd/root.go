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
		if (cmd.Flags().Lookup("c").Changed) {
			if (len(os.Args) == 4) {
				filename := os.Args[3]
				byteCount := CountBytes(filename)
				fmt.Println(byteCount, filename)
			} else {
				byteCount := CountBytes()
				fmt.Println(byteCount)
			}
		} else if (cmd.Flags().Lookup("w").Changed) {
			if (len(os.Args) == 4) {
				filename := os.Args[3]
				wordCount := CountWords(filename)
				fmt.Println(wordCount, filename)
			} else {
				wordCount := CountWords()
				fmt.Println(wordCount)
			}
		} else if (cmd.Flags().Lookup("l").Changed) {
			if (len(os.Args) == 4) {
				filename := os.Args[3]
				lineCount := CountLines(filename)
				fmt.Println(lineCount, filename)
			} else {
				lineCount := CountLines()
				fmt.Println(lineCount)
			}
		} else if (cmd.Flags().Lookup("m").Changed) {
			if (len(os.Args) == 4) {
				filename := os.Args[3]
				runeCount := CountRunes(filename)
				fmt.Println(runeCount, filename)
			} else {
				runeCount := CountRunes()
				fmt.Println(runeCount)
			}
		} else {
			if (len(os.Args) == 3) {
				filename := os.Args[2]

				wordCount := CountWords(filename)
				lineCount := CountLines(filename)
				runeCount := CountRunes(filename)
				fmt.Println(lineCount, wordCount, runeCount, filename)
			} else {
				wordCount := CountWords()
				lineCount := CountLines()
				runeCount := CountRunes()
				fmt.Println(lineCount, wordCount, runeCount)
			}

			
		}
	},
}

var (
	flag string
)

func Execute() {
	rootCmd.Flags().StringVarP(&flag, "c", "c", "", "count the number of bytes in a text file")
	rootCmd.Flags().Lookup("c").NoOptDefVal = " " 
	rootCmd.Flags().StringVarP(&flag, "w", "w", "", "count the number of words in a text file")
	rootCmd.Flags().Lookup("w").NoOptDefVal = " "
	rootCmd.Flags().StringVarP(&flag, "l", "l", "", "count the number of lines in a text file")
	rootCmd.Flags().Lookup("l").NoOptDefVal = " " 
	rootCmd.Flags().StringVarP(&flag, "m", "m", "", "count the number of characters in a text file")
	rootCmd.Flags().Lookup("m").NoOptDefVal = " "

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func CountWords(args ...string) int {
	var count int = 0

	if (len(args) == 0) {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			count++
		}
	} else if (len(args) == 1) {
		file, err := os.Open(args[0])

		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return 0
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			count++
		}
	}

	return count
}

func CountBytes(args ...string) int {
	var totalCount int

	if (len(args) == 0) {
		var continueCount bool = true

		for continueCount {
			buf := make([]byte, 1024)
			n, err := os.Stdin.Read(buf)
			if n == 0 {
				continueCount = false
			} else {
				totalCount += n
			
				if err != nil {
					fmt.Println("Error:", err)
					return 0
				}
			}
		}
	} else if (len(args) == 1) {
		file, err := os.Open(args[0])

		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return 0
		}

		defer file.Close()

		var continueCount bool = true

		for continueCount {
			buf := make([]byte, 1024)
			n, err := file.Read(buf)
			if n == 0 {
				continueCount = false
			} else {
				totalCount += n
			
				if err != nil {
					fmt.Println("Error:", err)
					return 0
				}
			}
		}
	}

	return totalCount
}

func CountLines(args ...string) int {
	var count int = 0

	if (len(args) == 0) {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			count++
		}	
		
	} else {
		file, err := os.Open(args[0])

		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return 0
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			count++
		}
	}
	
	return count
}

func CountRunes(args ...string) int {
	var count int = 0
	
	if (len(args) == 0) {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanRunes)

		for scanner.Scan() {
			count++
		}
		
	} else if (len(args) == 1) {
		file, err := os.Open(args[0])

		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return 0
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanRunes)

		for scanner.Scan() {
			count++
		}
	}

	return count
}