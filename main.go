package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Hello! :)")

	// Define command line flags
	uppercase := flag.Bool("uppercase", false, "Convert text to uppercase")
	lowercase := flag.Bool("lowercase", false, "Convert text to lowercase")
	sortLines := flag.Bool("sort", false, "Sort lines alphabetically")
	unique := flag.Bool("unique", false, "Remove duplicate lines")
	wordCount := flag.Bool("wordcount", false, "Count number of words")

	flag.Parse()

	// Check for input file
	if len(flag.Args()) < 1 {
		fmt.Println("Please specify an input file.")
		return
	}
	inputFile := flag.Args()[0]

	// Read the input file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	text := string(content)

	// Process text based on flags
	if *uppercase {
		text = strings.ToUpper(text)
	} else if *lowercase {
		text = strings.ToLower(text)
	}

	if *sortLines {
		lines := strings.Split(text, "\n")
		sort.Strings(lines)
		text = strings.Join(lines, "\n")
	}

	if *unique {
		lines := strings.Split(text, "\n")
		uniqueLines := make(map[string]struct{})
		var result []string

		for _, line := range lines {
			if _, exists := uniqueLines[line]; !exists {
				uniqueLines[line] = struct{}{}
				result = append(result, line)
			}
		}
		text = strings.Join(result, "\n")
	}

	if *wordCount {
		words := strings.Fields(text)
		fmt.Printf("Word count: %d\n", len(words))
		return // Exit after counting words
	}

	// Output the result
	fmt.Println(text)
}
