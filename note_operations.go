// note_operations.go
package main

import (
	"fmt"
	"strconv"
)

// Add a note
func addNote(notes []string) []string {
	fmt.Println("\nEnter the note text:")
	note := getUserInput()
	if note != "" {
		notes = append(notes, note)
	}
	return notes
}

// Delete a note
func deleteNote(notes []string) []string {
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
	notes = append(notes[:index], notes[index+1:]...)
	return notes
}

// Show notes
func showNotes(notes []string) {
	fmt.Println("\nNotes:")
	for i, note := range notes {
		fmt.Printf("%03d - %s\n", i+1, note)
	}
}
