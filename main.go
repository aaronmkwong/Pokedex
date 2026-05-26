// run from project root folder to test: go run .   
// terminate the program by pressing ctrl+c
// can run the CLIand tee the output (copies the stdout) to a new file called repl.log
// and .gitignore the log

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Print prompt without newline
		fmt.Print("Pokedex > ")

		// Read user input
		scanner.Scan()
		input := scanner.Text()

		// Clean input using helper function
		words := CleanInput(input)

		// Skip empty input
		if len(words) == 0 {
			continue
		}

		// First word is the command
		command := words[0]

		fmt.Printf("Your command was: %s\n", command)
	}
}