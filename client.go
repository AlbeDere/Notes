package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Print the help message
func printHelpMessage() {
	fmt.Println("Usage: ./notestool [COLLECTION_NAME]")
	fmt.Println("Manage a collection of notes.")
}

// Print the menu
func printMenu() {
	fmt.Println("\nSelect operation:")
	fmt.Println("1. Show notes.")
	fmt.Println("2. Add a note.")
	fmt.Println("3. Delete a note.")
	fmt.Println("4. Exit.")
}

// Get user input
func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
