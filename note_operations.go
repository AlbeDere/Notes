package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Note structure
type Note struct {
	Content   string
	Timestamp string
	Tags      []string
}

// Show notes
func showNotes(notes []Note) {
	fmt.Println("\nNotes:")
	for i, note := range notes {
		fmt.Printf("%03d - %s (Tags: %v, Timestamp: %s)\n", i+1, note.Content, note.Tags, note.Timestamp)
	}
}

// Add a note
func addNote() Note {
	fmt.Println("\nEnter the note text:")
	content := getUserInput()
	fmt.Println("Enter tags (comma-separated):")
	tags := strings.Split(getUserInput(), ",")

	return Note{
		Content:   content,
		Timestamp: time.Now().Format(time.RFC822),
		Tags:      tags,
	}
}

// Delete a note
func deleteNote(notes []Note) []Note {
	fmt.Println("\nEnter the number of note to remove or 0 to cancel:")
	input := getUserInput()
	noteNumber, err := strconv.Atoi(input)
	if err != nil || noteNumber < 0 || noteNumber > len(notes) {
		fmt.Println("Invalid note number.")
		return notes
	}

	if noteNumber == 0 {
		return notes
	}

	index := noteNumber - 1
	return append(notes[:index], notes[index+1:]...)
}
