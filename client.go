package main

import (
	"bufio"
	"os"
	"strings"
)

// Print the help message

// Print the menu

// Get user input
func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
