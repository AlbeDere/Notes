package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Load the collection from a file
func loadCollection(filePath string) ([]Note, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var notes []Note
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			parts := strings.SplitN(line, "|", 3)
			if len(parts) == 3 {
				timestamp := parts[0]
				tagsAndContent := strings.SplitN(parts[1], ";", 2)
				if len(tagsAndContent) == 2 {
					tags := strings.Split(tagsAndContent[0], ",")
					content := tagsAndContent[1]
					notes = append(notes, Note{
						Content:   content,
						Timestamp: timestamp,
						Tags:      tags,
					})
				}
			}
		}
	}

	return notes, scanner.Err()
}

// Save the collection to a file
func saveCollection(filePath string, notes []Note) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, note := range notes {
		tags := strings.Join(note.Tags, ",")
		line := fmt.Sprintf("%s|%s;%s", note.Timestamp, tags, note.Content)
		fmt.Fprintln(writer, line)
	}

	return writer.Flush()
}
