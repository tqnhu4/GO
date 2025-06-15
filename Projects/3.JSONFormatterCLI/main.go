package main

import (
	//"bytes"
	"fmt"
	"io"
	"os"
	"encoding/json"
)

func main() {
	// Determine input source: file or stdin
	var inputReader io.Reader
	var inputFilename string

	if len(os.Args) > 1 {
		// If an argument is provided, treat it as a filename
		inputFilename = os.Args[1]
		file, err := os.Open(inputFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", inputFilename, err)
			os.Exit(1)
		}
		defer file.Close()
		inputReader = file
	} else {
		// No argument provided, read from stdin
		inputReader = os.Stdin
	}

	// Read all content from the input source
	inputBytes, err := io.ReadAll(inputReader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Unmarshal the JSON data into a generic interface{}
	// This allows us to handle any valid JSON structure (object or array)
	var data interface{}
	err = json.Unmarshal(inputBytes, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Marshal the data back into JSON with indentation for pretty-printing
	// The prefix is for the beginning of each line, and indent for indentation.
	// We use an empty prefix and 4 spaces for indentation.
	formattedJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting JSON: %v\n", err)
		os.Exit(1)
	}

	// Print the formatted JSON to standard output
	fmt.Println(string(formattedJSON))
}