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
			filename := os.Args[3]
			byteCount := CountBytes(filename)
			fmt.Println(byteCount, filename)
		} else if (cmd.Flags().Lookup("w").Changed) {
			filename := os.Args[3]
			wordCount := CountWords(filename)
			fmt.Println(wordCount, filename)
		} else if (cmd.Flags().Lookup("l").Changed) {
			filename := os.Args[3]
			lineCount := CountLines(filename)
			fmt.Println(lineCount, filename)
		} else if (cmd.Flags().Lookup("m").Changed) {
			filename := os.Args[3]
			runeCount := CountRunes(filename)
			fmt.Println(runeCount, filename)
		} else {
			filename := os.Args[2]

			wordCount := CountWords(filename)
			lineCount := CountLines(filename)
			runeCount := CountRunes(filename)

			fmt.Println(lineCount, wordCount, runeCount, filename)
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

func CountWords(filename string) int {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return 0
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}

	return count
}

func CountBytes(filename string) int {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return 0
	}

	defer file.Close()

	var totalCount int
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

	return totalCount
}

func CountLines(filename string) int {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return 0
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		count++
	}

	return count
}

func CountRunes(filename string) int {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return 0
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	count := 0
	for scanner.Scan() {
		count++
	}

	return count
}