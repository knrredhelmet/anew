package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: anew <filename> [-o outputfile]")
		os.Exit(1)
	}

	filename := os.Args[1]
	var outputFileName string
	if len(os.Args) == 4 && os.Args[2] == "-o" {
		outputFileName = os.Args[3]
	}

	// Read existing lines from the file into a set
	existingLines := make(map[string]struct{})
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" { // Ignore blank lines
			existingLines[line] = struct{}{}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from file: %v\n", err)
		os.Exit(1)
	}

	// Create a scanner to read from stdin
	scanner = bufio.NewScanner(os.Stdin)
	newLines := []string{}

	// Read each line from stdin
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" { // Ignore blank lines
			if _, exists := existingLines[line]; !exists {
				newLines = append(newLines, line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	// If there are new unique lines, append them to the original file and print them
	if len(newLines) > 0 {
		// Open file for appending
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening file for writing: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		writer := bufio.NewWriter(file)

		// Check if the file ends with a newline
		fi, err := file.Stat()
		if err != nil {
			fmt.Printf("Error getting file stats: %v\n", err)
			os.Exit(1)
		}

		// Read the last byte to check if it ends with a newline
		if fi.Size() > 0 {
			buf := make([]byte, 1)
			_, err := file.Seek(-1, os.SEEK_END)
			if err != nil {
				fmt.Printf("Error seeking to the end of file: %v\n", err)
				os.Exit(1)
			}
			file.Read(buf)
			if buf[0] != '\n' {
				writer.WriteString("\n") // Add a newline if missing
			}
		}

		for _, line := range newLines {
			if _, err := writer.WriteString(line + "\n"); err != nil {
				fmt.Printf("Error writing to file: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(line) // Print new unique lines to stdout
		}

		writer.Flush() // Ensure all data is written to the file

		// If the -o option was provided, overwrite the output file with new lines
		if outputFileName != "" {
			outputFile, err := os.Create(outputFileName) // Always create/overwrite
			if err != nil {
				fmt.Printf("Error creating output file: %v\n", err)
				os.Exit(1)
			}
			defer outputFile.Close()

			for _, line := range newLines {
				if _, err := outputFile.WriteString(line + "\n"); err != nil {
					fmt.Printf("Error writing to output file: %v\n", err)
					os.Exit(1)
				}
			}
		}
	}
}
