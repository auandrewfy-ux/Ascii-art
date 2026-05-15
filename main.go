package main

import (
	""
	"os"
	"strings"
)

func main() {

	// Check if user provided an argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"text\"")
		return
	}

	input := os.Args[1]

	// Replace "\n" with real newline
	input = strings.ReplaceAll(input, "\\n", "\n")

	// Read the banner file
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading banner file")
		return
	}

	// Split banner file into lines
	lines := strings.Split(string(data), "\n")

	// Split input into lines (for handling \n)
	words := strings.Split(input, "\n")

	for w, word := range words {

		// If empty line, just print newline
		if word == "" {
			if w != len(words)-1 {
				fmt.Println()
			}
			continue
		}

		// Each ASCII character has height of 8 lines
		for i := 0; i < 8; i++ {

			for _, char := range word {

				// Find position of character in banner file
				index := (int(char) - 32) * 9

				// Print the correct line of the character
				fmt.Print(lines[index+i])
			}

			// Move to next row
			fmt.Println()
		}
	}
}
