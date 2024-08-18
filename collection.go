// collection.go
package main

import (
	"bufio"
	"os"
)

// Load the collection from a file
func loadCollection(filePath string) ([]string, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}

	return notes, scanner.Err()
}

// Save the collection to a file
func saveCollection(filePath string, notes []string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, note := range notes {
		_, err := writer.WriteString(note + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
