package main

import (
	"fmt"
	"os"
	"strings"

	// Import our local scanner package
	"github.com/AGX18/tiny-scanner/scanner"
)

func main() {
	fmt.Println("TINY Scanner Project")

	// 1. Get input file path from command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: tiny-scanner <input_file_path> [output_file_path]]")
		os.Exit(1)
	}
	inputFilePath := os.Args[1]

	// 2. Read the input file
	sourceCode, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Printf("Error reading file '%s': %v\n", inputFilePath, err)
		os.Exit(1)
	}

	// 3. Create a new scanner and scan the code
	fmt.Printf("Scanning file: %s\n", inputFilePath)
	myScanner := scanner.NewScanner(string(sourceCode))
	tokens := myScanner.ScanTokens()

	// 4. Create the output file
	// We'll name it "output.txt" and place it in the same directory.
	outputFilePath := "output.txt"
	if len(os.Args) > 2 && os.Args[2] != "" {
		outputFilePath = os.Args[2]
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	// 5. Write the tokens to the output file
	// Format: "tokenvalue, tokentype"
	for _, token := range tokens {
		if token.Type == scanner.EOF {
			break // Don't write the EOF token to the file
		}

		// Handle newlines in the input by writing their value as "newline"
		// This makes the output cleaner
		value := token.Value
		if strings.Contains(value, "\n") {
			value = "newline"
		}

		outputLine := fmt.Sprintf("%s, %s\n", value, token.Type.String())
		_, err := outputFile.WriteString(outputLine)
		if err != nil {
			fmt.Printf("Error writing to output file: %v\n", err)
		}
	}

	fmt.Printf("Successfully scanned %d tokens.\n", len(tokens))
	fmt.Printf("Output written to: %s\n", outputFilePath)
}
