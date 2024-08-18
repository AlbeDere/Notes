package main

import (
	"fmt"
	"os"
)

func main() {
	color := colors{
		Reset:     "\033[0m",
		Black:     "\033[30m",
		Red:       "\033[31m",
		Green:     "\033[32m",
		Yellow:    "\033[33m",
		Blue:      "\033[34m",
		Magenta:   "\033[35m",
		Cyan:      "\033[36m",
		Gray:      "\033[37m",
		White:     "\033[97m",
		Bold:      "\033[1m",
		Italic:    "\033[3m",
		Underline: "\033[4m",
		Invert:    "\033[7m"}

	if len(os.Args) != 2 || os.Args[1] == "help" {
		printHelpMessage()
		return
	}

	collectionName := os.Args[1]
	filePath := collectionName + ".txt"

	notes, err := loadCollection(filePath)
	if err != nil {
		fmt.Println("Error loading collection:", err)
		return
	}

	cycleProgram := true

	fmt.Print("\033[H\033[2J")
	fmt.Println(color.Bold + "Welcome to the Notes Tool! \n\n" + color.Reset)

	for cycleProgram {
		printMenu()
		choice := getUserInput()

		switch choice {
		case "1":
			fmt.Print("\033[H\033[2J")
			showNotes(notes)
		case "2":
			fmt.Print("\033[H\033[2J")
			note := addNote()
			notes = append(notes, note)
			saveCollection(filePath, notes)
		case "3":
			fmt.Print("\033[H\033[2J")
			notes = deleteNote(notes)
			saveCollection(filePath, notes)
		case "4":
			fmt.Println("Exiting the program.")
			cycleProgram = false
		default:
			fmt.Println("Invalid choice. Please select again.")
		}
	}
}
