package utils

import (
	"bufio"
	"fmt"
	"os"
)

func OpenFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	return file, nil
}

func CreateLineScanner(file *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}

// ReadFileLineByLine reads a file line by line and processes each line with the provided function.
func ReadFileLineByLine(filename string, processLine func(string)) error {
	// Open the file
	file, err := OpenFile(filename)
	if err != nil {
		return err
	}
	// Ensure the file is closed when we're done
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Create a scanner to read the file line by line
	scanner := CreateLineScanner(file)

	// Read and process each line
	for scanner.Scan() {
		processLine(scanner.Text())
	}

	// Handle any errors during scanning
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}
